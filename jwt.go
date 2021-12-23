package main

import (
	"crypto/rand"
	"fmt"
	"github.com/google/uuid"

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

func createToken(u *UserClaims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, u)
	signedToken, err := t.SignedString(keys[currentKid].key)
	if err != nil {
		return "", fmt.Errorf("error in creating token %w", err)
	}
	return signedToken, nil
}

func generateNewKey() error {
	newKey := make([]byte, 64)
	rand.Read(newKey)
	uid, err := uuid.NewUUID()
	if err != nil {
		return fmt.Errorf("error in generatingNewKey")
	}
	keys[uid.String()] = Key{
		key:     newKey,
		created: time.Now(),
	}
	currentKid = uid.String()
	return nil
}

type Key struct {
	key     []byte
	created time.Time
}

var currentKid = ""
var keys = map[string]Key{}

func parseToken(signedToken string, key []byte) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(signedToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("invalid signing Algorithm")
		}
		kid, ok := t.Header["kid"].(string)
		if !ok {
			return nil, fmt.Errorf("invalid key Id")
		}
		k, ok := keys[kid]
		if !ok {
			return nil, fmt.Errorf("invalid key Id")
		}
		return k, nil
	})
	if err != nil {
		return nil, fmt.Errorf("errro in parsetoken while paring token %w", err)
	}
	if !t.Valid {
		return nil, fmt.Errorf("Token is not valid")
	}
	claims := t.Claims.(*UserClaims)
	return claims, nil
}

func testJwt() {

}
