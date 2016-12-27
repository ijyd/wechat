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

type RespMsg struct {
	Code    int    `json:"ret_code"`
	ErrInfo string `json:"error_info"`
}

type AllocQRCode struct {
	RespMessage   RespMsg `json:"resp_msg"`
	DeviceID      string  `json:"deviceid"`
	QRCodeTicket  string  `json:"qrticket"`
	DeviceLicence string  `json:"devicelicence"`
}

type Device struct {
	ID                string `json:"id"`
	Mac               string `json:"mac"`
	ConnectProtocol   string `json:"connect_protocol"`
	AuthKey           string `json:"auth_key"`
	CloseStrategy     string `json:"close_strategy"`
	ConnStrategy      string `json:"conn_strategy"`
	CryptMethod       string `json:"crypt_method"`
	AuthVer           string `json:"auth_ver"`
	ManuMacPos        string `json:"manu_mac_pos"`
	SerMacPos         string `json:"ser_mac_pos"`
	BleSimpleProtocol string `json:"ble_simple_protocol"`
}

type DevAuthorizeReq struct {
	DeviceNum string `json:"device_num"`
	Devices   Device `json:"device_list"`
	OPType    string `json:"op_type"`
}

type DevAuthorizeResp struct {
}
