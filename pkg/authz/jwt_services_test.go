package authz

import (
	"crypto/rsa"
	"testing"
	"time"

	"github.com/gcarrenho/authz/test"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

var (
	jwtTestDefaultKey           *rsa.PublicKey
	jwtTestRSAPrivateKey        *rsa.PrivateKey
	jwtTestRsaKeyWOPublic       *rsa.PublicKey
	jwtTestunexpectedPrivateKey interface{}
)

func init() {
	// Load public keys
	jwtTestDefaultKey = test.LoadRSAPublicKeyFromDisk("test/sample_key.pub")
	jwtTestRsaKeyWOPublic = test.LoadRSAPublicKeyFromDisk("test/rsa_key_public.pub")

	// Load private keys
	jwtTestRSAPrivateKey = test.LoadRSAPrivateKeyFromDisk("test/sample_key.pem")
	jwtTestunexpectedPrivateKey = test.LoadFaildPrivateKeyFromDisk("test/unexpected_key.pem")

}

func TestValidateToken(t *testing.T) {
	tests := []struct {
		name          string
		claims        jwt.Claims
		publicKey     interface{}
		privateKey    interface{}
		valid         bool
		signingMethod jwt.SigningMethod // The method to sign the JWT token for test purpose
		err           string
	}{
		{
			name:          "Invalid public key crypto/rsa: verification error",
			publicKey:     jwtTestRsaKeyWOPublic,
			privateKey:    jwtTestRSAPrivateKey,
			claims:        jwt.MapClaims{"exp": time.Now().Add(10 * time.Minute), "authorized": true, "group_policy": "ManageDeviceSecurity"},
			err:           "crypto/rsa: verification error",
			signingMethod: jwt.SigningMethodRS256,
		},
		{
			name:          "Invalid signing method",
			publicKey:     jwtTestDefaultKey,
			privateKey:    jwtTestunexpectedPrivateKey,
			claims:        jwt.MapClaims{"exp": time.Now().Add(10 * time.Minute), "authorized": true, "group_policy": "ManageDeviceSecurity"},
			err:           "unexpected signing method: HS512",
			signingMethod: jwt.SigningMethodHS512,
		},
		{
			name:          "Invalid Policy",
			publicKey:     jwtTestDefaultKey,
			privateKey:    jwtTestRSAPrivateKey,
			claims:        jwt.MapClaims{"exp": time.Now().Add(10 * time.Minute), "authorized": true, "group_policy": "ManageDeviceSecurity"},
			err:           "invalid policy permission",
			signingMethod: jwt.SigningMethodRS256,
		},
		{
			name:          "Valid token",
			publicKey:     jwtTestDefaultKey,
			privateKey:    jwtTestRSAPrivateKey,
			claims:        jwt.MapClaims{"exp": time.Now().Add(10 * time.Minute), "authorized": true, "group_policy": "ManageDeviceSecurity,ReadDevices"},
			signingMethod: jwt.SigningMethodRS256,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, tokenString, _ := test.MakeSampleToken(tt.claims, tt.signingMethod, tt.privateKey)

			jwtSrv := NewJWTAuthService()

			err := jwtSrv.ValidateToken(tokenString, tt.publicKey)
			if err != nil {
				assert.Equal(t, tt.err, err.Error())
			}
		})
	}
}
