package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password []byte) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf(" hashPassword was not able to work: %w", err)
	}
	return hashedPassword, nil
}

func comparePassword(password, hashedPassword []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return fmt.Errorf("comparedPassword : invalid password %w", err)
	}
	return nil
}

func testcrpyt() {
	password := []byte("varun123")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	userinput := s.Bytes()
	hashedPassword, _ := hashPassword(password)
	err := comparePassword(userinput, hashedPassword)
	if err != nil {
		log.Panic("Invalid password")
	}
	fmt.Println("Logged in!")
}
