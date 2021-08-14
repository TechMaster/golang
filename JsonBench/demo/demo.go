package demo

import (
	goccy "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
)

type Person struct {
	Id      string
	Name    string
	Email   string
	Pass    string
	Roles   []string
	Age     int
	Enabled bool
}

var APerson Person
var People []Person
var StringJSON string

var Jsoniter jsoniter.API

func init() {
	Jsoniter = jsoniter.ConfigCompatibleWithStandardLibrary

	APerson = Person{
		Id:      "OX13",
		Name:    "Trinh Minh Cuong",
		Email:   "cuong@gmail.com",
		Pass:    "ok12323-",
		Roles:   []string{"admin", "manager", "editor"},
		Age:     46,
		Enabled: true,
	}

	People = []Person{
		{
			Id:      "OX13",
			Name:    "Trinh Minh Cuong",
			Email:   "cuong@gmail.com",
			Pass:    "ok12323-",
			Roles:   []string{"admin", "manager", "editor"},
			Age:     46,
			Enabled: true,
		},
		{
			Id:      "OX14",
			Name:    "Bui Van Hien",
			Email:   "hien@gmail.com",
			Pass:    "Alo alo",
			Roles:   []string{"teacher", "student", "editor"},
			Age:     24,
			Enabled: true,
		},
		{
			Id:      "OX15",
			Name:    "Nguyễn Trung Đức",
			Email:   "duc@gmail.com",
			Pass:    "Khong Ai Biet",
			Roles:   []string{"sysadmin", "student", "editor"},
			Age:     21,
			Enabled: false,
		},
		{
			Id:      "OX100",
			Name:    "Trinh Viet Linh",
			Email:   "linh@gmail.com",
			Pass:    "Cực kỳ cool",
			Roles:   []string{"web", "sale", "marketing"},
			Age:     26,
			Enabled: false,
		},
	}

	if buff, err := goccy.Marshal(APerson); err == nil {
		StringJSON = string(buff)
	}
}
