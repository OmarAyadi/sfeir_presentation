package dto

import (
	"time"
)

type ServerError struct {
	Error     string    `json:"error"`
	Timestamp time.Time `json:"timestamp"`
}

func NewServerError(err string) ServerError {
	return ServerError{
		Error:     err,
		Timestamp: time.Now(),
	}
}

func ErrorReturn(err error) ServerError {
	return NewServerError(err.Error())
}

type HealthReturn struct {
	Status string `json:"status,omitempty" example:"ok/down"`
}
