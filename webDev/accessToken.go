package webDev

import (
	"errors"
	"github.com/OOrangeeee/openwechat-sdk/webDev/util"
	"github.com/imroc/req"
	"strconv"
	"time"
)

type accessToken struct {
	token    string
	deadLine time.Time
}

type accessTokenTemp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

func (at *accessToken) getAccessToken(appid string, secret string) (string, error) {
	if at.isDead() {
		err := at.refreshAccessToken(appid, secret)
		if err != nil {
			return "", err
		}
	}
	return at.token, nil
}

func (at *accessToken) isDead() bool {
	if at.token == "" {
		return true
	}
	return time.Now().After(at.deadLine)
}

func (at *accessToken) refreshAccessToken(appid string, secret string) error {
	var tt accessTokenTemp
	param := req.Param{
		"grant_type": "client_credential",
		"appid":      appid,
		"secret":     secret,
	}
	url := accessTokenURL
	err := util.RequestGetJSONWithQueryParam(url, param, &tt)
	if err != nil {
		return err
	}
	if tt.ErrCode != 0 {
		return errors.New(strconv.Itoa(tt.ErrCode) + ":" + tt.ErrMsg)
	}
	at.token = tt.AccessToken
	at.deadLine = time.Now().Add(time.Second * time.Duration(tt.ExpiresIn-60))
	return nil
}
