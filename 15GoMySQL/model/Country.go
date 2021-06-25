package model

type Country struct {
	Code string `gorm:"primaryKey"`
	Name string
}
