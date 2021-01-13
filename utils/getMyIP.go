package utils

import (
	"errors"
	"github.com/golang/glog"
	"net/http"
	"strings"
)

const (
	getIPURL = "http://icanhazip.com"
)

var (
	ErrGetMyIPFail = errors.New("get my ip failed")
)

func GetMyIP() (string, error) {
	rCode, rBody, rError := SendRequest(http.MethodGet, getIPURL, nil, nil)
	if rError != nil {
		glog.Errorf("failed to get my ip, code: %d, body: %s, err: %v", rCode, rBody, rError)
		return "", ErrGetMyIPFail
	}

	return strings.TrimSpace(rBody), nil
}
