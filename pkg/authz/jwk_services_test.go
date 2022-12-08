package authz

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPublicKey(t *testing.T) {
	os.Setenv("ENVIRONMENT", "TEST")
	os.Setenv("JWKSURL", "https://pie.authz.wpp.api.hp.com/openid/v1/jwks.json")
	tests := []struct {
		name   string
		jwkurl string
		kid    string
		err    string
	}{
		{
			name:   "get public key Successful",
			jwkurl: os.Getenv("JWKSURL"),
			kid:    "authz-pie-1648037362",
		},
		{
			name:   "empty jwk resource",
			jwkurl: os.Getenv("JWKFAIL"),
			kid:    "authz-pie-1648037362",
			err:    "empty jwk resource",
		},
		{
			name:   "not found KID in token",
			jwkurl: os.Getenv("JWKSURL"),
			kid:    "authz-pie",
			err:    "not found KID in token",
		},
	}

	for _, tt := range tests {

		jwkSrv := NewJWKService(tt.jwkurl)
		_, err := jwkSrv.GetPublicKey(tt.kid)
		if err != nil {
			assert.Equal(t, tt.err, err.Error())
		}

	}
}
