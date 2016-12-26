package backend

const (
	//WeChatHookUnset wechat hook not set
	WeChatHookUnset = ""
	//WeChatHookIOT wechat hook iot backend
	WeChatHookIOT = "iot"
)

//Config contais all parameters required to start the wechat callback service
type Config struct {
	Type string

	OriginID       string
	AppID          string
	AppSecret      string
	Token          string
	EncodingAESKey string
}
