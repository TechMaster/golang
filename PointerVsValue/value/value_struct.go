package value

import (
	"github.com/TechMaster/eris"
	"github.com/go-playground/validator/v10"
	"github.com/segmentio/ksuid"
	"golang.org/x/crypto/bcrypt"
)

type AccountNew struct {
	Email    string `validate:"required,email"`
	FullName string `validate:"required,min=3,max=20"`
	Password string `validate:"required,min=4,max=15"`
}

type Account struct {
	Id       string //Uniqe Id
	Email    string
	FullName string
	HashPass string
}

type AccountRepo struct {
	accounts []Account //Che dấu dữ liệu tốt. Encapsulation in OOP
}

var AccountRepository AccountRepo

/* Hãy tận dụng hàm init này như là constructor cho Go package
 */
func init() {
	AccountRepository = AccountRepo{
		accounts: make([]Account, 0, 5),
	}

	_, _ = AccountRepository.Save(AccountNew{
		Email:    "tom@gmail.com",
		FullName: "Tom Sawyer",
		Password: "abc123-"})

	_, _ = AccountRepository.Save(AccountNew{
		Email:    "bob@gmail.com",
		FullName: "Bob Cun",
		Password: "abc123-"})

	_, _ = AccountRepository.Save(AccountNew{
		Email:    "anna@hotmail.com",
		FullName: "Anna Kournikova",
		Password: "abc123-"})
}

/* Khi viết hàm validate hãy lấy chính đối tượng làm receiver
 */
func (acc AccountNew) Validate() error {
	validate := validator.New()
	return validate.Struct(acc)
}

func (accRepo AccountRepo) GetAll() []Account {
	return accRepo.accounts
}

/*
Tại sao ở đây tôi dùng Pointer receiver?
Vì tôi tạo bản ghi và save accRepo.accounts !
*/
func (accRepo *AccountRepo) Save(acc AccountNew) (Id string, err error) {
	if err := acc.Validate(); err != nil {
		return "", err
	} else {
		saveAcc := Account{
			Id:       ksuid.New().String(),
			Email:    acc.Email,
			FullName: acc.FullName,
		}
		if hashPass, err := bcrypt.GenerateFromPassword([]byte(acc.Password), bcrypt.DefaultCost); err != nil {
			return "", eris.NewFromMsg(err, "Failed to hash password")
		} else {
			saveAcc.HashPass = string(hashPass)
		}

		accRepo.accounts = append(accRepo.accounts, saveAcc)

		return saveAcc.Id, nil
	}
}

//Chuyển từ trả về value sang pointer
func (accRepo *AccountRepo) GetById(id string) (acc Account, err error) {
	for _, account := range accRepo.accounts {
		if account.Id == id {
			return account, nil
		}
	}
	return Account{}, eris.New("Account not found")
}
