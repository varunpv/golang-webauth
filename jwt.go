package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims
	SessionId int64
}

func (uc UserClaims) Valid() error {
	if !uc.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("Token has expired")
	}
	if uc.SessionId == 0 {
		return fmt.Errorf("Invalid session ID")
	}
	return nil
}

func testJwt() {

}
