package connector

import (
	"context"
	"net/http"
	"net/url"
)

//Request is a handler for wechat message.
type Request interface {
	Hanlder(req *http.Request)
	Filter(req *http.Request) bool
}

//RequestHandler a handler for process request
type RequestHandler func(parameters map[string]string)

//WebHook implement whole wechat callback server require
type WebHook interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, parameters map[string]string, handler RequestHandler)
}

//Client implement request to wechat
type Client interface {
	Get(ctx context.Context, parameters map[string]string, out Object) error
	Post(ctx context.Context, parameters map[string]string, obj Object) error
}

//RESTClient is a common restclient in wechat
type RESTClient interface {
	URL() *url.URL
	ResourceName() string
	Client() *http.Client
}

//Interface contains whole wechat request and callback method
type Interface interface {
	WebHook
	Client
}

//AccessToken is wechat accesstoken approach
type AccessToken interface {
	Token() (token string, err error)
}

//Object all message need to registed as a Object
type Object interface {
	Type() string
}
