package domain

import "errors"

var (
	ErrNotFound       = errors.New("data not found")
	ErrConflict       = errors.New("data has been exist")
	ErrParamInput     = errors.New("param is not valid")
	ErrInternalServer = errors.New("internal server error")
	ErrForbidden      = errors.New("forbidden error")
)
