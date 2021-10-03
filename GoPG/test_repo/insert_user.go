package test_repo

import (
	"github.com/brianvoe/gofakeit/v6"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func gen_random_roles(numberOfRoles int) ([]int, []string) {
	int_roles := []int{}
	enum_roles := []string{}
	for i := 0; i < numberOfRoles; i++ {
		var role int
		for { //Loop cho đến khi tạo ra phần tử mới
			role = 1 + random.Intn(8)
			if !int_in_array(role, int_roles) {
				break
			}
		}

		int_roles = append(int_roles, role)
		enum_roles = append(enum_roles, ROLES[role])
	}
	return int_roles, enum_roles
}

func insert_user() error {
	var err error
	id, _ := gonanoid.New(8)
	int_roles, enum_roles := gen_random_roles(3)

	transaction, _ := DB.Begin()

	user := User{
		Id:         id,
		Name:       gofakeit.Name(),
		Int_roles:  int_roles,
		Enum_roles: enum_roles,
	}

	_, err = transaction.Model(&user).Insert()
	if !check_err(err, transaction) {
		return err
	}

	for _, role := range int_roles {
		user_role := User_Role{
			User_id: id,
			Role_id: role,
		}
		_, err = transaction.Model(&user_role).Insert()
		if !check_err(err, transaction) {
			return err
		}
	}

	err = transaction.Commit()
	return err
}
