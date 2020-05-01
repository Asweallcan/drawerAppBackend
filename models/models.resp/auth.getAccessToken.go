package models_resp

type AccessTokenResp struct {
	Access_token string `json:"access_token"`
	Expires_in   int `json:"expires_in"`
	Errcode      int `json:"errcode"`
	Errmsg       int `json:"errmsg"`
}
