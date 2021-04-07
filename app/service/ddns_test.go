package service

import (
	"flag"
	"fmt"
	"github.com/golang/glog"
	"github.com/skeyic/dnspod-ddns/utils"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func Test_ddnsService_GetCurrentIP(t *testing.T) {
	flag.Set("logtostderr", "true")
	flag.Set("v", "10")
	flag.Parse()

	svc := &DDNSService{
		loginToken:   "205809,de1aa624c7f624a13647678ccdb34ec6",
		domainID:     "85782217",
		recordID:     "736321772",
		recordLineID: "0",
		subDomain:    "www",
	}

	glog.V(4).Info(svc.GetCurrentIP())

}

func Test_ddnsService_DDNS(t *testing.T) {
	flag.Set("logtostderr", "true")
	flag.Set("v", "10")
	flag.Parse()

	svc := &DDNSService{
		loginToken:   "205809,de1aa624c7f624a13647678ccdb34ec6",
		domainID:     "85782217",
		recordID:     "736321772",
		recordLineID: "0",
		subDomain:    "www",
	}
	glog.V(4).Info(svc.DDNS("218.89.239.89"))
}

func TestGetMyIP2(t *testing.T) {

	req, err := http.NewRequest("GET", "https://api.my-ip.io/ip", nil)
	if err != nil {
		return
		// handle err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
		// handle err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("BODY: %s", body)

}

func TestGetMyIP4(t *testing.T) {
	rCode, rBody, rError := utils.SendRequest(http.MethodGet, "https://api.my-ip.io/ip", nil, nil)
	if rError != nil {
		glog.Errorf("failed to get my ip, code: %d, body: %s, err: %v", rCode, rBody, rError)
		return
	}

	fmt.Printf("IP: %s", strings.TrimSpace(rBody))
}
