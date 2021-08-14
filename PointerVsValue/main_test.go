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

var data = map[string]interface{}{
	"name":  "rock",
	"email": "rock@gmail.com",
	"pass":  "abc123",
}

func Benchmark_PassStructAsValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PassStructAsValue(acc)
	}
}

func Benchmark_PassStructAsPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PassStructAsPointer(&acc)
	}
}

//----------------------------------------------------------------
func Benchmark_PassMapAsValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PassMapAsValue(data)
	}
}

func Benchmark_PassMapAsPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PassMapAsPointer(&data)
	}
}
