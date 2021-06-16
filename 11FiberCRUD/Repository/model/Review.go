package model

type Review struct {
	Id      int64  `json:"Id"`
	BookId  int64  `json:"BookId"`
	Comment string `json:"comment"`
	Rating  int    `json:"rating"`
}
