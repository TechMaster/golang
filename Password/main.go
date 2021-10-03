package main

import (
	"fmt"
	"log"

	"github.com/alexedwards/argon2id"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "pa$$word"
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hash)
	fmt.Println(len(hash))

	// ComparePasswordAndHash performs a constant-time comparison between a
	// plain-text password and Argon2id hash, using the parameters and salt
	// contained in the hash. It returns true if they match, otherwise it returns
	// false.
	match, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Match: %v", match)

	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	hash = string(bytes)
	fmt.Println(hash)
	fmt.Println(len(hash))
}
