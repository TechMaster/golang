package main

import (
	"testing"
)

var acc = Account{
	id:    "OX13",
	name:  "Cuong",
	email: "cuong@gmail.com",
	pass:  "abc123",
}

func BenchmarkPassStructAsValue(b *testing.B) {

	for i := 0; i < b.N; i++ {
		PassStructAsValue(acc)
	}
}

func BenchmarkPassStructAsPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PassStructAsPointer(&acc)
	}
}
