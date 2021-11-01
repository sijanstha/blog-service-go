package jwt

// import (
// 	"fmt"
// 	"time"
// 	jwt "github.com/dgrijalva/jwt-go"
//   )

//   type JwtTokenGenerator interface {
// 	  GetToken(Payload) string
//   }

//   type JwtTokenValidator interface{
// 	  ValidateToken(string) Payload
//   }
  
//   func GetJWT() (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)
  
// 	claims := token.Claims.(jwt.MapClaims)
  
// 	claims["authorized"] = true
// 	claims["client"] = "Krissanawat"
// 	claims["aud"] = "billing.jwtgo.io"
// 	claims["iss"] = "jwtgo.io"
// 	claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
  
// 	tokenString, err := token.SignedString(mySigningKey)
  
// 	if err != nil {
// 	  fmt.Errorf("Something Went Wrong: %s", err.Error())
// 	  return "", err
// 	}
  
// 	return tokenString, nil
//   }
 

//   type Payload struct {
// 	  id string
// 	  username string
// 	  role string
//   }
  