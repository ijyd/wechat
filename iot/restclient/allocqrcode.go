package restclient

import (
	"net/http"
	"net/url"
	"reflect"

	"github.com/golang/glog"
	"github.com/ijyd/wechat/iot/api"
)

func init() {
	url, err := url.Parse("https://api.weixin.qq.com/device/getqrcode")
	if err != nil {
		glog.Fatalf("parse raw url error %v\r\n", err)
	}

	resource := reflect.TypeOf(&api.AllocQRCode{}).String()
	devAuth := &allocQRCode{
		base:         url,
		resourceName: resource,
		client:       iotClient,
	}

	RegisterClient(devAuth)
}

type allocQRCode struct {
	base         *url.URL
	resourceName string
	client       *http.Client
}

func (d allocQRCode) URL() *url.URL {
	return d.base
}

func (d allocQRCode) ResourceName() string {
	return d.resourceName
}

func (d allocQRCode) Client() *http.Client {
	return d.client
}
