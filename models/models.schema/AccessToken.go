package models_schema

import "github.com/jinzhu/gorm"

type (
	AccessToken struct {
		gorm.Model
		Value string
		ExpiresIn int `gorm:"column:expires_in"`
	}
)

func (AccessToken) TableName() string {
	return "access_tokens"
}
