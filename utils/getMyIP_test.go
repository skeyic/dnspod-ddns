package utils

import (
	"flag"
	"github.com/golang/glog"
	"testing"
)

func TestGetMyIP(t *testing.T) {
	flag.Set("logtostderr", "true")
	flag.Set("v", "10")
	flag.Parse()

	glog.V(4).Info(GetMyIP())
}
