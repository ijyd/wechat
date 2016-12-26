package connector

import "net/http"

//Write write response
func Write(statusCode int, w http.ResponseWriter, req *http.Request, obj Object) {

	switch obj.Type() {
	case "error":
		err := obj.(*Error)
		errStr := err.Error()
		w.Write([]byte(errStr))
	}

}
