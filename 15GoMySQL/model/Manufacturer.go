package model

type Manufacturer struct {
	ID          uint
	Name        string
	CountryCode string `json:"-"`
	Country     *Country
}

/*
Table manufacturers {
  ID int [pk, increment]
  name varchar
  country_code varchar[2] [ref: > countries.code]
}
*/
