package jwt

type TokenGenerator interface {
	GetToken(Payload) (string, error)
}

type TokenValidator interface {
	ValidateToken(string) (Payload, error)
}

type Payload struct {
	Id    string
	Email string
	Role  string
}
