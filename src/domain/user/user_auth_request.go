package user

import stringutils "github.com/blog-service/src/utils/string"

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	UserDetails UserDomain `json:"userDetails"`
	Token       string     `json:"token"`
}

func (loginRequest *UserLoginRequest) Validate() error {
	if stringutils.IsEmptyOrNull(loginRequest.Email) {
		return ErrMissingEmail
	}
	if !stringutils.IsValidEmail(loginRequest.Email) {
		return ErrInvalidEmail
	}
	if stringutils.IsEmptyOrNull(loginRequest.Password) {
		return ErrMissingPassword
	}
	return nil
}
