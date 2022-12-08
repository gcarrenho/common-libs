package test

import (
	"crypto/rsa"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func LoadFaildPrivateKeyFromDisk(location string) interface{} {
	keyData, e := os.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	return keyData
}

func LoadRSAPrivateKeyFromDisk(location string) *rsa.PrivateKey {
	keyData, e := os.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

func LoadRSAPublicKeyFromDisk(location string) *rsa.PublicKey {
	keyData, e := os.ReadFile(location)
	if e != nil {
		panic(e.Error())
	}
	key, e := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if e != nil {
		panic(e.Error())
	}
	return key
}

// MakeSampleToken creates and returns a encoded JWT token that has been signed with the specified cryptographic key.
func MakeSampleToken(claims jwt.Claims, signingMethod jwt.SigningMethod, privateKey interface{}) (*jwt.Token, string, error) {
	token := jwt.NewWithClaims(signingMethod, claims)

	tokenString, e := token.SignedString(privateKey)
	if e != nil {
		panic(e.Error())
	}

	return token, tokenString, nil
}
