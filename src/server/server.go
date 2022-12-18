package server

type ServerPort interface {
	StartApplication(addr string)
}
