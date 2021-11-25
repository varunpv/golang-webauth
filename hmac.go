package main

import (
	"bufio"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var key = make([]byte, 64)

func testhmac() {
	rand.Read(key)
	fmt.Println("Enter your username")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	input := s.Bytes()
	signedMessaged, _ := signMessage(input)

	for {
		fmt.Println("This is your Token. Keep it safe! -> ", signedMessaged)
		fmt.Println("Did you note it down? YES or NO?")
		s.Scan()
		isnoted := strings.ToLower(s.Text())
		if isnoted == "yes" {
			break
		}
	}
	fmt.Println("okay let me test you.... what is the token? ")
	s.Scan()
	inputeToken := getbyteslice(s.Text())
	if ok, _ := validateMessage(inputeToken, input); ok {
		fmt.Println("Nice !")
	} else {
		fmt.Println(" _|_")
	}

}

func getbyteslice(s string) []byte {
	var result []byte
	stringnum := strings.Fields(s)
	for _, is := range stringnum {
		i, err := strconv.Atoi(is)
		if err == nil {
			result = append(result, byte(i))
		}
	}
	return result
}

func signMessage(message []byte) ([]byte, error) {
	h := hmac.New(sha512.New, key)
	_, err := h.Write(message)
	if err != nil {
		return nil, fmt.Errorf("error in sighMessage function : %w", err)
	}
	return h.Sum(nil), nil
}

func validateMessage(signedMessaged, message []byte) (bool, error) {
	correctSign, err := signMessage(message)
	if err != nil {
		return false, fmt.Errorf("error in validating message %w", err)
	}
	isvalid := hmac.Equal(correctSign, signedMessaged)
	return isvalid, nil
}
