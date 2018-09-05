package comm

import (
	"bytes"
	"net/http"

	"github.com/json-iterator/go"
)

func NewClient(appKey string) *Client {
	return &Client{appKey}
}

type Client struct {
	AppKey string
}

func (c *Client) Send(r Request) *Response {
	buf := GetBuffer()
	defer PutBuffer(buf)
	err := r.MarshalURL(c, buf)
	resp := &Response{err: err}
	if err != nil {
		return resp
	}
	resp.httpResp, resp.err = http.DefaultClient.Get(string(buf.Bytes()))
	if resp.err != nil {
		return resp
	}
	if resp.httpResp.StatusCode != http.StatusOK {
		resp.err = ErrUnknown
	}
	return resp
}

type Request interface {
	MarshalURL(*Client, *bytes.Buffer) error
}

type Response struct {
	err      error
	httpResp *http.Response
}

func (this *Response) Unmarshal(v interface{}) error {
	if this.err != nil {
		return this.err
	}
	defer this.httpResp.Body.Close()
	return jsoniter.ConfigFastest.NewDecoder(this.httpResp.Body).Decode(v)
}
