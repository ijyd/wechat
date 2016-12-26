package factory

import (
	"fmt"

	"github.com/ijyd/wechat/connector"
	"github.com/ijyd/wechat/connector/backend"
)

//Create  appropriate processing backends based on configuration parameters
func Create(c backend.Config) (connector.Interface, error) {

	switch c.Type {
	case backend.WeChatHookIOT, backend.WeChatHookUnset:
		return newIOT(c)
	}

	return nil, fmt.Errorf("not reach here")
}
