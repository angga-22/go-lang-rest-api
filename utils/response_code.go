package utils

import "net/http"

func GenerateCodeFromError(err error) int{
	switch err {
	case ERR_CREATE_DATABASE:
		return http.StatusInternalServerError
	case ERR_EMPTY_DATA_IN_DATABASE:
		return http.StatusOK
	case ERR_GET_DATA_IN_DATABASE:
		return http.StatusInternalServerError
	}
	return http.StatusOK
}