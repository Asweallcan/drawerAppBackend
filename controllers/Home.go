package controllers

import (
	"drawerBackend/api"
	"drawerBackend/constants"
	"drawerBackend/models"
	models_resp "drawerBackend/models/models.resp"
	"drawerBackend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(ctx *gin.Context) {
	var (
		resBody  interface{} = nil
		resError models.Error
	)

	accessToken, err := utils.GetAccessToken()
	resError = err

	if utils.HasNoError(resError) {
		temp, err := utils.Request(
			"Post",
			fmt.Sprintf(api.CLOUD_FUNCTION, accessToken, constants.CLOUD_ENV, "main"),
			gin.H{
				"POSTBODY": gin.H{
					"$url":    "test",
					"hahahah": "wocao",
				}})

		if err != nil {
			resError = utils.Error(constants.WX_CLOUD_FUNCTION_FAILED, constants.COMMON_ERROR_MESSAGE)
		} else {
			var resp models_resp.CloudFunctionResp
			_ = utils.MapToStruct(temp, &resp)
			resBody = resp
		}
	}

	ctx.JSON(http.StatusOK, utils.Response(resBody, resError))
}
