package utils

import (
	"net/http"

	"restful.api.e-commerce.golang/domain"
)

func GetHttpStatus(err error) int {
	switch err {
	case domain.ErrInternalServer:
		return http.StatusInternalServerError

	case domain.ErrConflict:
		return http.StatusConflict

	case domain.ErrParamInput:
		return http.StatusUnprocessableEntity

	case domain.ErrNotFound:
		return http.StatusNotFound

	default:
		return http.StatusInternalServerError
	}
}
