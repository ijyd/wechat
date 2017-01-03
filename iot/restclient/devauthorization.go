package restclient

import (
	"net/http"
	"net/url"
	"reflect"

	"github.com/golang/glog"
	"github.com/ijyd/wechat/iot/api"
)

func init() {
	url, err := url.Parse("https://api.weixin.qq.com/device/authorize_device")
	if err != nil {
		glog.Fatalf("parse raw url error %v\r\n", err)
	}
	glog.V(5).Infof("init rest client\r\n")

	resource := reflect.TypeOf(&api.DevAuthorizeReq{}).String()
	devAuth := &devAuthorization{
		base:         url,
		resourceName: resource,
		client:       iotClient,
	}

	RegisterClient(devAuth)
}

type devAuthorization struct {
	base         *url.URL
	resourceName string
	client       *http.Client
}

func (d devAuthorization) URL() *url.URL {
	return d.base
}

func (d devAuthorization) ResourceName() string {
	return d.resourceName
}

func (d devAuthorization) Client() *http.Client {
	return d.client
}
