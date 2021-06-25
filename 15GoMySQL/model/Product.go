package model

type Product struct {
	ID             uint
	Name           string
	Description    string
	Price          uint
	Madein         string
	Country        *Country `gorm:"foreignKey:Madein"`
	ManufacturerID uint
	Manufacturer   *Manufacturer
	CategoryID     uint
	Category       *Category
}

/*
Table products {
  ID int [pk, increment]
  name varchar [not null]
  description varchar [not null]
  madein varchar(2) [ref: > countries.code]
  price int
  manufacturer_id int [ref: > manufacturers.ID]
  category_id int [ref: > categories.ID]
}
*/
