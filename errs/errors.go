package errs

import "net/http"

type AppError struct {
	// json tag used to omit properties which are empty from displaying when the AppError error is encoded
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

// MessageOnly takes an AppError error and returns an AppError error with only the message set - this allows us to encode
// the error message in json without the code (the code is already set by WriteHeader(err.Code))
func (e AppError) MessageOnly() *AppError {
	return &AppError{Message: e.Message}
}

func NotFoundErr(message string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: message}
}

func UnexpectedErr(message string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: message}
}

func ValidationErr(message string) *AppError {
	return &AppError{Code: http.StatusUnprocessableEntity, Message: message}
}

func AuthorizationErr(message string) *AppError {
	return &AppError{Code: http.StatusUnauthorized, Message: message}
}

/*func BadRequestErr(message string) *AppError {
	return &AppError{Code: http.StatusBadRequest, Message: message}
}*/
