package utils

import (
	"drawerBackend/api"
	"drawerBackend/constants"
	"drawerBackend/models"
	models_resp "drawerBackend/models/models.resp"
	"drawerBackend/models/models.schema"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func GetAccessToken() (string, models.Error) {
	var (
		accessToken models_schema.AccessToken
		ret         string
		retError    models.Error
	)

	db, err := ConnectDB()

	if err != nil {
		retError = Error(constants.DB_ERROR, constants.COMMON_ERROR_MESSAGE)
	} else {
		defer db.Close()

		if !db.HasTable(models_schema.AccessToken{}) {
			db.CreateTable(models_schema.AccessToken{})
		}

		db.First(&accessToken)
		if accessToken.UpdatedAt.Unix()+int64(accessToken.ExpiresIn) > time.Now().Unix() {
			ret = accessToken.Value
		} else {
			if accessToken == (models_schema.AccessToken{}) {
				accessToken.ID = 1
				db.Create(&accessToken)
			}

			temp, err := Request("Post", fmt.Sprintf(api.ACCESS_TOKEN, constants.AppId, constants.AppSecret), nil)
			if err != nil {
				retError = Error(constants.WX_FAILED, constants.TOKEN_FAIL_MESSAGE)
			} else {
				var resp models_resp.AccessTokenResp
				_ = MapToStruct(temp, &resp)
				if resp.Errcode == constants.WX_SUCCESS {
					db.Model(&accessToken).Updates(gin.H{"value": resp.Access_token, "expires_in": resp.Expires_in})
					ret = resp.Access_token
				} else {
					retError = Error(constants.WX_FAILED, constants.TOKEN_FAIL_MESSAGE)
				}
			}
		}
	}

	return ret, retError
}
