package authz

import (
	"context"
	"errors"

	"github.com/lestrrat-go/jwx/jwk"
)

type JWKServices interface {
	GetPublicKey(kid string) (interface{}, error)
}

type jwkServices struct {
	set jwk.Set
}

// NewJWKService initialize a new jwkServices
func NewJWKService(jwkURL string) *jwkServices {
	return &jwkServices{
		set: getKeys(jwkURL),
	}
}

// GetPublicKey take a kid and return a public key
func (jwksrv jwkServices) GetPublicKey(kid string) (interface{}, error) {
	if jwksrv.set == nil {
		return nil, errors.New("empty jwk resource")
	}

	key, err := lookupKeyID(kid, jwksrv.set)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwk.PublicRawKeyOf(key)
	if err != nil {
		return nil, err
	}

	return publicKey, err
}

// getKeys do a request to the jwkURL parameter and return the keys
func getKeys(jwkURL string) jwk.Set {
	keys, err := jwk.Fetch(context.Background(), jwkURL)
	if err != nil {
		return nil
	}
	return keys
}

// lookupKeyID looking the match kid in the set and return the key if this match exist or return error if dont exist
func lookupKeyID(kid string, set jwk.Set) (jwk.Key, error) {
	key, exists := set.LookupKeyID(kid)
	if !exists {
		return nil, errors.New("not found KID in token")
	}
	key.Set("alg", "RS256")

	return key, nil
}
