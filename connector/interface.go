package connector

import "net/http"

//Request is a handler for wechat message.
type Request interface {
	Hanlder(req *http.Request)
	Filter(req *http.Request) bool
}

//RequestHandler a handler for process request
type RequestHandler func(parameters map[string]string)

//Hook implement whole wechat callback server reqire
type Hook interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, parameters map[string]string, handler RequestHandler)
}

//Client implement request to wechat
type Client interface {
	Get()
	Post()
}

//Interface contains whole wechat request and callback method
type Interface interface {
	Hook
	Client
}

//Object all message need to registed as a Object
type Object interface {
	Type() string
}
