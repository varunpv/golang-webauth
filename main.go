package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := []byte("123")
	hashPass, err := hashPassword(pass)
	if err != nil {
		log.Panic(err)
	}
	err = comparePassword(hashPass, pass)
	if err != nil {
		log.Println("Invalid password")
	}
	fmt.Println("Logged in!")

}

func hashPassword(password []byte) ([]byte, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return []byte{}, fmt.Errorf("Unable to hash password from bcrypt.GenerateFromPassword, %w", err)
	}
	return hashedPassword, nil
}

func comparePassword(hashedPassword, password []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return fmt.Errorf("Unsuccessful comparasion: %w", err)
	}
	return nil

}
