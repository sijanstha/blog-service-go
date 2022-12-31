package mapper

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

func ToGrpcErrorCode(httpStatusCode int) codes.Code {
	switch httpStatusCode {
	case http.StatusBadRequest:
		return codes.InvalidArgument

	case http.StatusInternalServerError:
		return codes.Internal

	case http.StatusNotFound:
		return codes.NotFound

	case http.StatusUnauthorized:
		return codes.Unauthenticated

	default:
		return codes.OK
	}
}
