package resp

type CloudFunctionResp struct {
	Errcode   int
	eErrmsg    string
	Resp_data interface{}
}
