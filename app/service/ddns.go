package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/golang/glog"
	"github.com/skeyic/dnspod-ddns/config"
	"github.com/skeyic/dnspod-ddns/utils"
	"net/http"
	"time"
)

const (
	recordListURL = "https://dnsapi.cn/Record.List"
	ddnsURL       = "https://dnsapi.cn/Record.Ddns"
)

var (
	ErrGetCurrentIP = errors.New("get current ip failed")
	ErrDDNS         = errors.New("ddns failed")
)

type DDNSService struct {
	loginToken   string
	domainID     string
	recordID     string
	recordLineID string
	subDomain    string

	previousIP string
}

func NewDDNSServiceFromConfig() *DDNSService {
	return &DDNSService{
		loginToken:   config.Config.DDNS.LoginToken,
		domainID:     config.Config.DDNS.DomainID,
		recordID:     config.Config.DDNS.RecordID,
		recordLineID: config.Config.DDNS.RecordLineID,
		subDomain:    config.Config.DDNS.SubDomain,
	}
}

type dnsRecord struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

// curl -X POST https://dnsapi.cn/Record.List -d 'login_token=205809,de1aa624c7f624a13647678ccdb34ec6&domain=xiaxuanli.com&format=json' -s | python -mjson.tool
func (d *DDNSService) GetCurrentIP() (string, error) {
	requestBody := fmt.Sprintf("login_token=%s&domain_id=%s&sub_domain=%s&format=json", d.loginToken, d.domainID, d.subDomain)
	rCode, rBody, rError := utils.SendRequest(http.MethodPost, recordListURL, bytes.NewBufferString(requestBody), utils.NewXWwwFormUrlencodedSendOptions())
	if rCode != http.StatusOK || rError != nil {
		glog.Errorf("failed to get current ip, code: %d, body: %s, err: %v", rCode, rBody, rError)
		return "", ErrGetCurrentIP
	}
	records, _, _, err := jsonparser.Get([]byte(rBody), "records")
	if err != nil {
		glog.Errorf("failed to unmarshal current ip, body: %s, err: %v", rBody, err)
		return "", ErrGetCurrentIP
	}

	var (
		dnsRecords []*dnsRecord
	)

	err = json.Unmarshal(records, &dnsRecords)
	if err != nil {
		glog.Errorf("failed to unmarshal records, err: %v", err)
		return "", ErrGetCurrentIP
	}

	if len(dnsRecords) != 1 {
		glog.Errorf("failed to get records, err: %v", err)
		return "", ErrGetCurrentIP
	}

	return dnsRecords[0].Value, nil
}

// curl -X POST https://dnsapi.cn/Record.Ddns -d 'login_token=205809,de1aa624c7f624a13647678ccdb34ec6&
// format=json&domain_id=85782217&record_id=736321772&record_line_id=0&sub_domain=www&value=1.1.1.1'
func (d *DDNSService) DDNS(ip string) error {
	requestBody := fmt.Sprintf("login_token=%s&domain_id=%s&record_id=%s&record_line_id=%s&sub_domain=%s&value=%s&format=json",
		d.loginToken, d.domainID, d.recordID, d.recordLineID, d.subDomain, ip)
	rCode, rBody, rError := utils.SendRequest(http.MethodPost, ddnsURL, bytes.NewBufferString(requestBody), utils.NewXWwwFormUrlencodedSendOptions())
	if rCode != http.StatusOK || rError != nil {
		glog.Errorf("failed to ddns, code: %d, body: %s, err: %v", rCode, rBody, rError)
		return ErrDDNS
	}
	d.previousIP = ip
	glog.V(4).Infof("ddns to %s successfully", ip)
	return nil
}

func (d *DDNSService) Process() {
	var processFunc = func() {
		glog.V(4).Infof("processing")
		myIP, err := utils.GetMyIP()
		if err != nil {
			glog.Errorf("skip this round, failed to get my ip: %v", err)
			return
		}
		if myIP == d.previousIP {
			glog.V(4).Infof("skip this round, IP does not change")
			return
		}
		if d.previousIP == "" {
			currentIP, err := d.GetCurrentIP()
			if err != nil {
				glog.V(4).Infof("skip this round, failed to get current IP")
				return
			}
			d.previousIP = currentIP
			if currentIP == myIP {
				glog.V(4).Infof("skip this round, IP does not change")
				return
			}
		}
		err = d.DDNS(myIP)
		if err != nil {
			glog.Error("failed to DDNS this round, need to change the ip to %s", myIP)
			return
		}
		glog.V(4).Infof("PROCESS change IP to %s", myIP)
	}

	var (
		ticker    = time.NewTicker(config.Config.TimeInterval * time.Second)
		startChan = make(chan struct{}, 1)
	)

	go func() {
		time.Sleep(30 * time.Second)
		glog.V(4).Infof("kick up at the beginning")
		startChan <- struct{}{}
	}()

	for {
		select {
		case <-ticker.C:
			processFunc()
		case <-startChan:
			processFunc()
		}
	}
}
