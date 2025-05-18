package domain

import "fmt"

type CustomError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func (ce *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", ce.Code, ce.Message)
}

func BadRequestError(message string, errors interface{}) *CustomError {
	return &CustomError{
		Code:    400,
		Message: message,
		Errors:  errors,
	}
}

func NotFoundError(message string, errors interface{}) *CustomError {
	return &CustomError{
		Code:    404,
		Message: message,
		Errors:  errors,
	}
}

func InternalServerError(message string, errors interface{}) *CustomError {
	return &CustomError{
		Code:    500,
		Message: message,
		Errors:  errors,
	}
}

func UnauthorizedError(message string, errors interface{}) *CustomError {
	return &CustomError{
		Code:    401,
		Message: message,
		Errors:  errors,
	}
}