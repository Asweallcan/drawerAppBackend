package utils

import (
	"drawerBackend/api"
	"drawerBackend/constants"
	"drawerBackend/models"
	"drawerBackend/models/resp"
	"drawerBackend/models/schema"
	"fmt"
	"net/http"
	"time"
)

func GetAccessToken() (string, interface{}) {
	db, err := ConnectDB()
	var accessToken schema.AccessToken
	if err == nil {

		defer db.Close()

		if !db.HasTable(&schema.AccessToken{}) {
			db.CreateTable(&schema.AccessToken{})
		}

		db.First(&accessToken)

		if accessToken != (schema.AccessToken{}) {
			if accessToken.UpdatedAt.Unix()+int64(accessToken.ExpiresIn) > time.Now().Unix() {
				return accessToken.Value, nil
			}
		} else {
			accessToken.ID = 1
			db.Create(&accessToken)
		}

		res, err := http.Post(fmt.Sprintf(api.ACCESS_TOKEN, constants.AppId, constants.AppSecret), "application/json", nil)

		if err == nil {
			accessTokenResp := &resp.AccessTokenResp{}
			err := UnmarshalResp(res.Body, &accessTokenResp)

			if accessTokenResp.Errcode == constants.WX_SUCCESS && err == nil {
				db.Model(&accessToken).Updates(map[string]interface{}{"value": accessTokenResp.Access_token, "expires_in": accessTokenResp.Expires_in})
				return accessTokenResp.Access_token, nil
			}
		}

	}

	return accessToken.Value, models.Error{
		Code:    constants.WX_FAILED,
		Message: constants.TOKEN_FAIL_MESSAGE,
	}
}
