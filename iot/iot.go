package iot

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/golang/glog"
	"github.com/ijyd/wechat/connector"
	"github.com/ijyd/wechat/connector/client"
	"github.com/ijyd/wechat/iot/api"
	"github.com/ijyd/wechat/iot/restclient"
)

//IOT implement wechat iot callback service
type IOT struct {
	Token          string
	EncodingAESKey string
	tokenHandler   connector.AccessToken
}

//NewIOT Create  iot wechat hook
func NewIOT(token, aeskey string, accessToken connector.AccessToken) (connector.Interface, error) {
	glog.V(5).Infof("init  NewIOT\r\n")

	return &IOT{
		Token:          token,
		EncodingAESKey: aeskey,
		tokenHandler:   accessToken,
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

	token, err := i.tokenHandler.Token()
	if err != nil {
		return err
	}
	glog.V(5).Infof("Get method parameters(%v) token(%v)  out.kind(%v)",
		parameters, token, reflect.TypeOf(out))

	glog.V(5).Infof("restclient %v\r\n", restclient.ClientFunc)
	restC, ok := restclient.ClientFunc[reflect.TypeOf(out).String()]
	if !ok {
		glog.Errorf("not found client func with type(%v)\r\n", reflect.TypeOf(out))
		return fmt.Errorf("not found client func")
	}

	req := client.NewRequest(nil, http.MethodGet, restC.URL())
	req.SetParam("access_token", token)
	for k, v := range parameters {
		req.SetParam(k, v)
	}

	err = req.Request(func(req *http.Request, resp *http.Response) error {
		if resp.StatusCode >= 300 {
			glog.Errorf("response error code %v\r\n", resp.StatusCode)
			return fmt.Errorf("response error %v", resp.StatusCode)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			glog.Errorf("read body error %v", err)
			return err
		}
		resp.Body.Close()
		glog.V(5).Infof("resp body\r\n", string(body))

		return nil
	})

	return err
}

//Post implement Get method for client interface
func (i *IOT) Post(ctx context.Context, parameters map[string]string, obj connector.Object) error {
	token, err := i.tokenHandler.Token()
	if err != nil {
		return err
	}
	glog.V(5).Infof("Get method parameters(%v) token(%v)  out.kind(%v)",
		parameters, token, reflect.TypeOf(obj))

	glog.V(5).Infof("restclient %v\r\n", restclient.ClientFunc)
	restC, ok := restclient.ClientFunc[reflect.TypeOf(obj).String()]
	if !ok {
		glog.Errorf("not found client func with type(%v)\r\n", reflect.TypeOf(obj))
		return fmt.Errorf("not found client func")
	}

	req := client.NewRequest(nil, http.MethodPost, restC.URL())
	req.SetParam("access_token", token)
	for k, v := range parameters {
		req.SetParam(k, v)
	}

	body, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	req.Body(body)
	err = req.Request(func(req *http.Request, resp *http.Response) error {
		if resp.StatusCode >= 300 {
			glog.Errorf("response error code %v\r\n", resp.StatusCode)
			return fmt.Errorf("response error %v", resp.StatusCode)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			glog.Errorf("read body error %v", err)
			return err
		}
		resp.Body.Close()
		glog.V(5).Infof("resp body\r\n", string(body))

		return nil
	})

	return err
}
