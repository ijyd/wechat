package client

import (
	"net/http"
	"net/url"

	"github.com/golang/glog"
	"github.com/ijyd/wechat/iot"
)

func init() {
	url, err := url.Parse("https://api.weixin.qq.com/device/authorize_device")
	if err != nil {
		glog.Fatalf("parse raw url error %v\r\n", err)
	}

	devAuth := &devAuthorization{
		base:         url,
		resourceName: "",
		client:       iotClient,
	}

	iot.RegisterClient(devAuth)
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
