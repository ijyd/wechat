package accesstoken

import "github.com/chanxuehong/wechat.v2/mp/core"

type wechatv2token struct {
	srv core.AccessTokenServer
}

func (t wechatv2token) Token() (token string, err error) {
	token, err = t.srv.Token()
	return
}
