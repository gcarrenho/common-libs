package authz

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

type JWTServices interface {
	ValidateToken(headerToken string, publicKey interface{}) error
}

type jwtServices struct {
}

// NewJWTAuthService initialize the jwtServices
func NewJWTAuthService() JWTServices {
	return &jwtServices{}
}

// ValidateToken take the headerToken and the publicKey and check if this token is valid
func (service *jwtServices) ValidateToken(headerToken string, publicKey interface{}) error {
	token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			if alg, fail := token.Header["alg"].(string); fail {
				return nil, errors.New("unexpected signing method: " + alg)
			}
		}
		return publicKey, nil
	})

	if err != nil {
		return err
	}

	if !validatePolicy(token) {
		return errors.New("invalid policy permission")
	}

	return nil
}

// validatePolicy take a jwt token and check if it contain DeviceEventSubscriber policy
func validatePolicy(token *jwt.Token) bool {
	claims := token.Claims.(jwt.MapClaims)
	policies := claims["group_policy"].(string)
	policy := strings.Contains(policies, "ReadDevices")

	return policy
}
