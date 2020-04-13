package utils

import (
	"drawerBackend/models"
)

func Response(data interface{}, error interface{}) models.Response {
	if error != nil {
		return models.Response{
			Error: models.Error{
				Code:    error.(models.Error).Code,
				Message: error.(models.Error).Message,
			},
			Success: false,
			Data:    nil,
		}
	}
	return models.Response{
		Error:   models.Error{},
		Success: true,
		Data:    data,
	}
}
