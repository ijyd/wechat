package backend

const (
	//WeChatUnset wechat hook not set
	WeChatUnset = ""
	//WeChatIOT wechat hook iot backend
	WeChatIOT = "iot"
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
