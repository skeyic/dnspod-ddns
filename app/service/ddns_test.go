package service

import (
	"flag"
	"github.com/golang/glog"
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
	glog.V(4).Info(svc.DDNS("125.70.215.133"))
}
