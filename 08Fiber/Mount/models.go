package main

type Book struct {
	Title  string  `json:"name"`
	Author string  `json:"author"`
	Rating float32 `json:"rating"`
}

type Book2 struct {
	Id      int
	Title   string   `json:"name"`
	Authors []Author `json:"authors"`
	Rating  float32  `json:"rating"`
}

type Author struct {
	FullName string
	Country  string
}
