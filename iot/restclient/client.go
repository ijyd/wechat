package restclient

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/ijyd/wechat/connector"
)

var iotClient = http.DefaultClient

//ClientFunc a map store RESTClient interface. struct kind as key
var ClientFunc = make(map[string]connector.RESTClient)

//RegisterClient registed client interface
func RegisterClient(c connector.RESTClient) {
	glog.V(5).Infof("install %v \r\n", c.ResourceName())
	ClientFunc[c.ResourceName()] = c
}
