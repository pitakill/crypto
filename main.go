package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//Hash implements root.Hash
type Hash struct{}

//Generate a salted hash for the input string
func (c *Hash) Generate(s string) (string, error) {
	saltedBytes := []byte(s)
	hashedBytes, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hash := string(hashedBytes[:])
	return hash, nil
}

//Compare string to generated hash
func (c *Hash) Compare(hash string, s string) error {
	incoming := []byte(s)
	existing := []byte(hash)

	return bcrypt.CompareHashAndPassword(existing, incoming)
}

func main() {
	h := Hash{}

	data, err := h.Generate("hola mundo")
	if err != nil {
		panic(err)
	}

	if err := h.Compare(data, "hola mundo"); err != nil {
		panic(err)
	}

	fmt.Println("Everything is smooth!")
}
