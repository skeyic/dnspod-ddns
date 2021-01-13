package utils

import (
	"bytes"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"time"
)

type SendOptions struct {
	contentType string
}

func NewXWwwFormUrlencodedSendOptions() *SendOptions {
	return &SendOptions{contentType: "application/x-www-form-urlencoded"}
}

func NewJsonSendOptions() *SendOptions {
	return &SendOptions{contentType: "application/json; charset=utf-8"}
}

func (s *SendOptions) ContentType() string {
	return s.contentType
}

// SendRequest ...
func SendRequest(method string, uri string, body *bytes.Buffer, options *SendOptions) (int, string, error) {
	var (
		responseBody string
	)

	client := &http.Client{}
	client.Timeout = time.Minute * 3

	glog.V(6).Info(method)
	glog.V(6).Info(uri)
	glog.V(6).Info(body.String())

	if body == nil {
		body = bytes.NewBuffer([]byte("{}"))
	}
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		glog.Fatalf("http.NewRequest() failed with '%s'\n", err)
	}

	if options == nil {
		options = NewJsonSendOptions()
	}
	req.Header.Set("Content-Type", options.ContentType())
	resp, err := client.Do(req)
	if err != nil {
		glog.Warningf("client.Do() failed with '%s'\n", err)
		return http.StatusBadRequest, "", err
	}
	glog.V(6).Info(resp.StatusCode)
	glog.V(6).Infof("RESP: %+v", resp)

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	responseBody = string(bodyBytes)

	return resp.StatusCode, responseBody, nil
}
