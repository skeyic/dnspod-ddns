package main

import (
	"github.com/golang/glog"
	"github.com/skeyic/dnspod-ddns/app/service"
)

func main() {
	svc := service.NewDDNSServiceFromConfig()
	glog.V(4).Infof("DDNS Service is starting")
	go svc.Process()

	<-make(chan struct{}, 1)
}
