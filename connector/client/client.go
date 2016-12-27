package connector

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/golang/glog"
)

// Request allows for building up a request to a server in a chained fashion.
// Any errors are stored until the end of your call, so you only have to
// check once.
type Request struct {
	// required
	client *http.Client
	verb   string

	baseURL *url.URL
	timeout time.Duration

	params  url.Values
	headers http.Header

	// output
	err  error
	body io.Reader

	// The constructed request and the response
	req  *http.Request
	resp *http.Response
}

// NewRequest creates a new request helper object for accessing runtime.Objects on a server.
func NewRequest(client *http.Client, verb string, baseURL *url.URL) *Request {
	r := &Request{
		client:  client,
		verb:    verb,
		baseURL: baseURL,
	}
	r.SetHeader("Accept", "*/*")

	return r
}

//SetParam set parameter for url
func (r *Request) SetParam(paramName, value string) *Request {
	if r.params == nil {
		r.params = make(url.Values)
	}
	r.params[paramName] = append(r.params[paramName], value)
	return r
}

//SetHeader set header for http request
func (r *Request) SetHeader(key, value string) *Request {
	if r.headers == nil {
		r.headers = http.Header{}
	}
	r.headers.Set(key, value)
	return r
}

// URL returns the current working URL.
func (r *Request) URL() *url.URL {

	finalURL := &url.URL{}
	if r.baseURL != nil {
		*finalURL = *r.baseURL
	}

	query := url.Values{}
	for key, values := range r.params {
		for _, value := range values {
			query.Add(key, value)
		}
	}

	finalURL.RawQuery = query.Encode()
	return finalURL
}

// request implement send request to remote server and extract response
func (r *Request) request(fn func(*http.Request, *http.Response)) error {
	//Metrics for total request latency
	start := time.Now()

	if r.err != nil {
		glog.V(4).Infof("Error in request: %v", r.err)
		return r.err
	}

	client := r.client
	if client == nil {
		client = http.DefaultClient
	}

	url := r.URL().String()
	req, err := http.NewRequest(r.verb, url, r.body)
	if err != nil {
		return err
	}
	req.Header = r.headers

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	done := func() bool {
		// Ensure the response body is fully read and closed
		// before we reconnect, so that we reuse the same TCP
		// connection.
		defer func() {
			const maxBodySlurpSize = 2 << 10
			if resp.ContentLength <= maxBodySlurpSize {
				io.Copy(ioutil.Discard, &io.LimitedReader{R: resp.Body, N: maxBodySlurpSize})
			}
			resp.Body.Close()
		}()

		fn(req, resp)
		return true
	}()

	glog.V(5).Infof("request end result(%v) Spend time (%vs)", done, time.Now().Second()-start.Second())
	return nil
}
