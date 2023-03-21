package utils

import "errors"

var (
	ERR_CREATE_DATABASE = errors.New("Error ketika input data ke database")
	ERR_EMPTY_DATA_IN_DATABASE = errors.New("Data masih kosong")
	ERR_GET_DATA_IN_DATABASE = errors.New("Error ketika get data di database")

)
