package union

import (
	"net/http"

	"github.com/ijyd/wechat/wechathook"
)

type unionRequestHandler struct {
	Handlers    []wechathook.Request
	FailOnError bool
}

//New create a unionRequestHandle without fail on error
func New(requestHandlers ...wechathook.Request) wechathook.Request {
	if len(requestHandlers) == 1 {
		return requestHandlers[0]
	}
	return &unionRequestHandler{
		Handlers:    requestHandlers,
		FailOnError: false,
	}
}

//NewFailOnError create a unionRequestHandle with fail on error
func NewFailOnError(requestHandlers ...wechathook.Request) wechathook.Request {
	if len(requestHandlers) == 1 {
		return requestHandlers[0]
	}
	return &unionRequestHandler{
		Handlers:    requestHandlers,
		FailOnError: true,
	}
}

func (request *unionRequestHandler) Hanlder(req *http.Request) {

}

func (request *unionRequestHandler) Filter(req *http.Request) bool {
	return true
}
