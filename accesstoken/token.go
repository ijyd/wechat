package accesstoken

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/ijyd/wechat/connector"
)

//WeChatAccessToken implement AccessToken interface. hold current token for service
type WeChatAccessToken struct {
	appID     string
	appSecret string
	expire    time.Duration

	tokenMux sync.RWMutex
	token    string
}

type token struct {
	connector.WeChatError
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

const (
	defaultExpire  time.Duration = time.Duration(2 * time.Hour)
	wechatTokenURL               = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential"
)

//Token get current token
func (t *WeChatAccessToken) Token() (token string, err error) {
	t.tokenMux.RLock()
	token = t.token
	t.tokenMux.RUnlock()

	return token, nil
}

func (t *WeChatAccessToken) tokenCenter() {
	aheadTime := time.Duration(15 * time.Minute)
	if t.expire < aheadTime {
		glog.Fatalf("invalid token expire %v, must greater 15(minute)\r\n", t.expire)
	}

	refresh := t.expire - time.Duration(15*time.Minute)
	httpClient := http.DefaultClient
	url := fmt.Sprintf("%s&appid=%s&secret=%s", wechatTokenURL, t.appID, t.appSecret)

	for {
		select {
		case <-time.After(refresh):

			token, err := requestToken(httpClient, url)
			if err != nil {
				glog.Errorf("request token error %v\r\n", err)
				refresh = t.tokenRefreshTime(0)
				continue
			}

			if token.ErrCode != 0 {
				glog.Errorf("wechat result errorcode(%v) errmsg(%v)\r\n", token.ErrCode, token.ErrMsg)
				refresh = t.tokenRefreshTime(0)
				continue
			}

			glog.V(5).Infof("refresh token new(%+v)\r\n", token)

			t.tokenMux.Lock()
			t.token = token.Token
			t.tokenMux.Unlock()

			//determine next refresh time
			refresh = t.tokenRefreshTime(token.ExpiresIn)
		}
	}
}

func (t *WeChatAccessToken) tokenRefreshTime(expire int64) time.Duration {
	refresh := t.expire - time.Duration(15*time.Minute)

	if expire == 0 {
		rand.Seed(time.Now().UTC().UnixNano())
		waitTime := rand.Intn(120) + 30
		refresh = time.Duration(waitTime)
	} else if expire < int64(time.Duration(5*time.Minute)) {
		glog.Warningf("get a small expire token(%v), check your default token expire\r\n", expire)
		refresh = time.Duration(3 * time.Minute)
	} else if expire > int64(time.Duration(2*time.Hour)) {
		glog.Warningf("get a large expire token(%v), check your default token expire\r\n", expire)
		refresh = t.expire - time.Duration(15*time.Minute)
	} else {
		refresh = time.Duration(expire - int64(time.Duration(5*time.Minute)))
	}

	return refresh
}

func requestToken(c *http.Client, url string) (*token, error) {
	httpResp, err := c.Get(url)
	if err != nil {
		return nil, fmt.Errorf("refresh wechat token error %v\r\n", err)
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("refresh wechat token http resp error code %v\r\n", httpResp.StatusCode)
	}

	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, fmt.Errorf("read http body error %v\r\n", err)
	}

	tokenResp := token{}
	err = json.Unmarshal(body, tokenResp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal http body error %v\r\n", err)
	}

	return &tokenResp, nil
}
