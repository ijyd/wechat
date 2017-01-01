package accesstoken

import (
	"sync"

	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/ijyd/wechat/connector"
)

//WeChatToken a global handler for wechat token
var WeChatToken connector.AccessToken

//Startup create wechat access token and start internal goroutine
func Startup(appID, appSecret string) {

	var once sync.Once
	onceFunc := func() {
		token := &WeChatAccessToken{
			appID:     appID,
			appSecret: appSecret,
			expire:    defaultExpire,
		}

		WeChatToken = token
		go token.tokenCenter()
	}

	once.Do(onceFunc)
}

//StartupWechatV2 create wechat v2 token
func StartupWechatV2(srv core.AccessTokenServer) {

	var once sync.Once
	onceFunc := func() {
		token := &wechatv2token{
			srv: srv,
		}

		WeChatToken = token
	}

	once.Do(onceFunc)
}
