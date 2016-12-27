package iot

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
	"github.com/ijyd/wechat/connector"
	"github.com/ijyd/wechat/iot/api"
)

//IOT implement wechat iot callback service
type IOT struct {
	Token          string
	EncodingAESKey string
}

var clientFunc map[string]connector.RESTClient

//RegisterClient registed client interface
func RegisterClient(c connector.RESTClient) {
	clientFunc[c.ResourceName()] = c
}

//NewIOT Create  iot wechat hook
func NewIOT(token, aeskey string) (connector.Interface, error) {
	return &IOT{
		Token:          token,
		EncodingAESKey: aeskey,
	}, nil
}

func (i *IOT) ServeHTTP(w http.ResponseWriter, r *http.Request, parameters map[string]string, handler connector.RequestHandler) {
	method := r.Method

	switch method {
	case http.MethodPost:
		glog.V(5).Infof("Get parameters %+v\r\n", parameters)
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			err := connector.NewInvalidBodyError(err.Error(), "")
			connector.Write(400, w, r, err)
			return
		}
		defer r.Body.Close()

		errObj := i.processAsyncMessage(body)
		if err != nil {
			connector.Write(500, w, r, errObj)
			return
		}

	default:
		err := connector.NewMethodNotAllowError("not allow method", "")
		connector.Write(415, w, r, err)
		return
	}

	connector.Write(200, w, r, nil)

}

func (i *IOT) processAsyncMessage(body []byte) *connector.Error {
	meta := api.AsyncMeta{}
	err := json.Unmarshal(body, meta)
	if err != nil {
		return connector.NewInvalidBodyError(err.Error(), "")
	}

	glog.V(5).Infof("Raw messsage %v \r\n", string(body))

	switch meta.MsgType {
	case api.MsgTypeNotify:
		notify := api.Notify{}
		err := json.Unmarshal(body, &notify)
		if err != nil {
			glog.Errorf("unmarshal json failure %v\r\n", err)
			return connector.NewInvalidBodyError(err.Error(), "")
		}
		glog.V(5).Infof("Get  message %+v）\r\n", notify)
	case api.MsgTypeGetResp:
		getResp := api.GetResp{}
		err := json.Unmarshal(body, &getResp)
		if err != nil {
			glog.Errorf("unmarshal json failure %v\r\n", err)
			return connector.NewInvalidBodyError(err.Error(), "")
		}
		glog.V(5).Infof("Get  message %+v\r\n", getResp)
	case api.MsgTypeSetResp:
		setResp := api.SetResp{}
		err := json.Unmarshal(body, &setResp)
		if err != nil {
			glog.Errorf("unmarshal json failure %v\r\n", err)
			return connector.NewInvalidBodyError(err.Error(), "")
		}
		glog.V(5).Infof("Get  message %+v\r\n", setResp)
	default:
		return connector.NewInvalidBodyError(fmt.Errorf("invalid message type(%v)", meta.MsgType).Error(), "")
	}

	return nil
}

//Get implement Get method for client interface
func (i *IOT) Get(ctx context.Context, parameters map[string]string, out connector.Object) error {
	return nil
}

//Post implement Get method for client interface
func (i *IOT) Post(ctx context.Context, parameters map[string]string, obj connector.Object) error {
	return nil
}
