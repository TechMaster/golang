package model

type Category struct {
	ID       uint
	Name     string
	ParentID uint
	Parent   *Category
}
