package utils

import (
	"drawerBackend/models"
)

func Response(data interface{}, err models.Error) models.Response {
	if !HasNoError(err) {
		return models.Response{
			Error: models.Error{
				Code:    err.Code,
				Message: err.Message,
			},
			Success: false,
			Data:    nil,
		}
	}

	return models.Response{
		Error:   err,
		Success: true,
		Data:    data,
	}
}
