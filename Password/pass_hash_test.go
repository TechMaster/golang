package main

import (
	"testing"

	"github.com/alexedwards/argon2id"
	"golang.org/x/crypto/bcrypt"
)

var password = "pa$$word"

func Benchmark_Argon2id_Hash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = argon2id.CreateHash(password, argon2id.DefaultParams)
	}
}

func Benchmark_BCrypt_Hash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = bcrypt.GenerateFromPassword([]byte(password), 10)
	}
}
