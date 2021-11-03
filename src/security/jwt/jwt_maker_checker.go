package jwt

import (
	"os"
	"time"

	"github.com/blog-service/src/utils/logger"
	stringutils "github.com/blog-service/src/utils/string"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	token_signing_key = os.Getenv("token_signing_key")
	token_expiry_time = stringutils.ParseInteger(os.Getenv("token_expiry_time"))
)

type JwtTokenService struct{}

func (service *JwtTokenService) GetToken(request Payload) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = request.Email
	claims["id"] = request.Id
	claims["role"] = request.Role
	claims["aud"] = "blog.service.io"
	claims["iss"] = "blog.service.io"
	claims["exp"] = time.Now().Add(time.Duration(token_expiry_time) * time.Minute).Unix()

	tokenString, err := token.SignedString([]byte(token_signing_key))

	if err != nil {
		logger.Error("error while creating token", err)
		return "", err
	}

	return tokenString, nil
}

func (service *JwtTokenService) ValidateToken(token string) (*Payload, error) {
	return nil, nil
}
