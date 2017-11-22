package main

import "net/http"

type apiError struct {
	Error   error
	Message string
	Code    int
}

func internalServerError(err error) *apiError {
	return &apiError{
		Error:   err,
		Message: "Internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func notFound(err error) *apiError {
	return &apiError{
		Error:   err,
		Message: "Not Found",
		Code:    http.StatusNotFound,
	}
}

func badRequest(err error) *apiError {
	return &apiError{
		Error:   err,
		Message: "Bad Request",
		Code:    http.StatusBadRequest,
	}
}
