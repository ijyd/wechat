package connector

const (
	//WeChatErrCodeOK result is ok
	WeChatErrCodeOK = 0
	//WeChatErrCodeInvalidCred invalid appid or appsecret
	WeChatErrCodeInvalidCred = 40001
	//WeChatErrCodeAccessTokenExpired token expire
	WeChatErrCodeAccessTokenExpired = 42001 // access_token 过期错误码(maybe!!!)
)

//WeChatError wechat common error struct
type WeChatError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
