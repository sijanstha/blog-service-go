package jwt

import (
	"os"
	"testing"

	"github.com/blog-service/src/utils/logger"
	stringutils "github.com/blog-service/src/utils/string"
)

func TestGetToken(t *testing.T) {
	os.Setenv("token_signing_key", "test")
	os.Setenv("token_expiry_time", "10")
	service := &JwtTokenService{}
	token, err := service.GetToken(Payload{
		Id:    "1",
		Email: "test@testing.com",
		Role:  "ROLE_USER",
	})

	if err != nil {
		t.Errorf("failed to generate token: %v", err.Error())
	}

	if stringutils.IsEmptyOrNull(token) {
		t.Errorf("empty token")
	}

	logger.Info(token)
}
