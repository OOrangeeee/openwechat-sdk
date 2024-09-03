package webDev

import (
	"errors"
	"github.com/OOrangeeee/openwechat-sdk/webDev/util"
	"github.com/imroc/req"
	"strconv"
	"time"
)

type jsapiTicket struct {
	ticket   string
	deadLine time.Time
}

type jsapiTicketTemp struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int64  `json:"expires_in"`
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
}

func (jt *jsapiTicket) getJsapiTicket(accessToken string) (string, error) {
	if jt.isDead() {
		err := jt.refreshJsapiTicket(accessToken)
		if err != nil {
			return "", err
		}
	}
	return jt.ticket, nil
}

func (jt *jsapiTicket) isDead() bool {
	if jt.ticket == "" {
		return true
	}
	return time.Now().After(jt.deadLine)
}

func (jt *jsapiTicket) refreshJsapiTicket(accessToken string) error {
	var tt jsapiTicketTemp
	param := req.Param{
		"access_token": accessToken,
		"type":         "jsapi",
	}
	url := jsapiTicketURL
	err := util.RequestGetJSONWithQueryParam(url, param, &tt)
	if err != nil {
		return err
	}
	if tt.ErrCode != 0 {
		return errors.New(strconv.Itoa(tt.ErrCode) + ":" + tt.ErrMsg)
	}
	jt.ticket = tt.Ticket
	jt.deadLine = time.Now().Add(time.Second * time.Duration(tt.ExpiresIn-60))
	return nil
}
