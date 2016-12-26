package factory

import (
	"github.com/ijyd/wechat/connector"
	"github.com/ijyd/wechat/connector/backend"
	"github.com/ijyd/wechat/iot"
)

func newIOT(c backend.Config) (connector.Interface, error) {
	return iot.NewIOTWeChatHook(c.Token, c.EncodingAESKey)
}
