package models

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	error := fmt.Sprintf("Code: %d, Message: %s", e.Code, e.Message)
	logrus.Error(error)
	return error
}
