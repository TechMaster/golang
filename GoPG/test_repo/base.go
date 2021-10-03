package test_repo

import (
	"math/rand"
	"time"

	"github.com/go-pg/pg/v10"
)

const (
	ADMIN      = 1
	STUDENT    = 2
	TRAINER    = 3
	SALE       = 4
	EMPLOYER   = 5
	AUTHOR     = 6
	EDITOR     = 7 //edit bài, soạn page, làm công việc digital marketing
	MAINTAINER = 8 //quản trị hệ thống, gánh bớt việc cho Admin, back up dữ liệu. Sửa đổi profile,role user, ngoại trừ role ROOT và Admin
)

var (
	//Phần tử đầu tiên bỏ qua
	ROLES  = []string{"", "ADMIN", "STUDENT", "TRAINER", "SALE", "EMPLOYER", "AUTHOR", "EDITOR", "MAINTAINER"}
	DB     *pg.DB
	random *rand.Rand
)

type User struct {
	tableName  struct{} `pg:"demo.users"`
	Id         string   `pg:"id,pk"`
	Name       string
	Int_roles  []int    `pg:"int_roles,array"`
	Enum_roles []string `pg:"enum_roles,array"`
}

type User_Role struct {
	tableName struct{} `pg:"demo.user_role"`
	User_id   string
	Role_id   int
}

func init() {
	DB = pg.Connect(&pg.Options{
		Addr:     "192.168.1.9:5432",
		User:     "postgres",
		Password: "123",
		Database: "dvdrental",
	})
	s1 := rand.NewSource(time.Now().UnixNano())
	random = rand.New(s1)
}
