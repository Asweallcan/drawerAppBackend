package controllers

import (
	"drawerBackend/api"
	"drawerBackend/constants"
	"drawerBackend/models/resp"
	"drawerBackend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(ctx *gin.Context) {
	var resBody resp.CloudFunctionResp
	accessToken, err := utils.GetAccessToken()
	if err == nil {
		reqBody, _ := utils.RequestBody(gin.H{
			"POSTBODY": gin.H{
				"$url": "test",
			},
		})
		res, _ := http.Post(fmt.Sprintf(api.CLOUD_FUNCTION, accessToken, constants.CLOUD_ENV, "main"), "application/json", reqBody)
		_ = utils.UnmarshalResp(res.Body, &resBody)
		ctx.JSON(http.StatusOK, utils.Response(resBody, nil))
	} else {
		ctx.JSON(http.StatusInternalServerError, utils.Response(nil, err))
	}
}
