package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type UserClaims struct {
	jwt.StandardClaims
	SessionId int64
}

func (uc UserClaims) Valid() error {
	if !uc.VerifyExpiresAt(time.Now().Unix(), true) {
		return fmt.Errorf("token has expired")
	}
	if uc.SessionId == 0 {
		return fmt.Errorf("invalid session ID")
	}
	return nil
}

func createToken(u *UserClaims, key []byte) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, u)
	signedtoken, err := t.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("error in creating token %w", err)
	}
	return signedtoken, nil
}

func parseToken(signedToken string,key []byte) (*UserClaims, error) {
	t,err:= jwt.ParseWithClaims(signedToken,&UserClaims{},func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("invalid signing Algorithm")
		}
		return key,nil
	})
	if err != nil {
		return nil, fmt.Errorf("errro in parsetoken while paring token %w",err)
	}
	if !t.Valid {
		return nil, fmt.Errorf("Token is not valid")
	}
	claims := t.Claims.(*UserClaims)
	return claims,nil
}

func testJwt() {
	key := make([]byte, 64)
	rand.Read(key)

}
