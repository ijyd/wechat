package api

//Response contais a response message for wechat
type Response struct {
	Code      string `json:"error_code"`
	Message   string `json:"error_msg"`
	MessageID string `json:"msg_id,omitempty"`
}

//Type indicate  object type
func (*Response) Type() string {
	return "Response"
}

const (
	//MsgTypeNotify indicate notify message from wechat
	MsgTypeNotify = "notify"
	//MsgTypeGetResp indicate query response from wechat
	MsgTypeGetResp = "get_resp"
	//MsgTypeSetResp indicate set device response from wechat
	MsgTypeSetResp = "set_resp"
)

type AsyncMeta struct {
	MsgID   int64  `json:"msg_id"`
	MsgType string `json:"msg_type"`
}

type SetResp struct {
	AsyncMeta

	AsyErrorCode int    `json:"asy_error_code"`
	AsyErrorMsg  string `json:"asy_error_msg"`
}

//Type indicate  object type
func (*SetResp) Type() string {
	return "SetResp"
}

type GetResp struct {
	AsyncMeta

	AsyErrorCode int    `json:"asy_error_code"`
	AsyErrorMsg  string `json:"asy_error_msg"`
}

//Type indicate  object type
func (*GetResp) Type() string {
	return "GetResp"
}

//Notify IoT notify message
type Notify struct {
	AsyncMeta

	DeviceType string  `json:"device_type"`
	DeviceID   string  `json:"device_id"`
	Service    Service `json:"services"`
	Data       string  `json:"data"`
}

//Type indicate  object type
func (*Notify) Type() string {
	return "Notify"
}

type Service struct {
	OperationStatus string `json:"operation_status"`
	Status          int8   `json:"status"`
}
