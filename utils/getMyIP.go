package utils

import (
	"errors"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	getIPURLOld = "http://icanhazip.com"
	getIPURL    = "https://api.my-ip.io/ip"
	getIPURLV3  = "https://api.ipify.org"
	getIPURLV4  = "https://ipinfo.io/ip"
)

var (
	ErrGetMyIPFail = errors.New("get my ip failed")
)

func GetMyIPOld() (string, error) {
	rCode, rBody, rError := SendRequest(http.MethodGet, getIPURLOld, nil, nil)
	if rError != nil {
		glog.Errorf("failed to get my ip, code: %d, body: %s, err: %v", rCode, rBody, rError)
		return "", ErrGetMyIPFail
	}

	return strings.TrimSpace(rBody), nil
}

func GetMyIPOld2() (string, error) {
	rCode, rBody, rError := SendRequest(http.MethodGet, getIPURL, nil, nil)
	if rError != nil {
		glog.Errorf("failed to get my ip, code: %d, body: %s, err: %v", rCode, rBody, rError)
		return "", ErrGetMyIPFail
	}

	return strings.TrimSpace(rBody), nil
}

func GetMyIPOld3() (string, error) {
	rCode, rBody, rError := SendRequest(http.MethodGet, getIPURLV3, nil, nil)
	if rError != nil {
		glog.Errorf("failed to get my ip, code: %d, body: %s, err: %v", rCode, rBody, rError)
		return "", ErrGetMyIPFail
	}

	return strings.TrimSpace(rBody), nil
}

// GetMyIP V4
func GetMyIP() (string, error) {
	req, err := http.NewRequest("GET", "https://ipinfo.io/ip", nil)
	if err != nil {
		// handle err
	}
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7,ja;q=0.6")
	req.Header.Set("Cache-Control", "max-age=0")
	req.Header.Set("Priority", "u=0, i")
	req.Header.Set("^Sec-Ch-Ua", "^^Chromium^^;v=^^128^^, ^^Not;A=Brand^^;v=^^24^^, ^^Google")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("^Sec-Ch-Ua-Platform", "^^Windows^^^")
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "none")
	req.Header.Set("Sec-Fetch-User", "?1")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return strings.TrimSpace(string(bodyBytes)), nil
}
