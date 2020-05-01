package utils

import "drawerBackend/models"

func Error(code int, message string) models.Error {
	return models.Error{
		Code:    code,
		Message: message,
	}
}

func HasNoError(err models.Error) bool {
	return err == models.Error{}
}
