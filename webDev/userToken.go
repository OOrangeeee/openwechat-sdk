package webDev

import (
	"errors"
	"github.com/OOrangeeee/openwechat-sdk/webDev/util"
	"github.com/imroc/req"
	"strconv"
	"time"
)

type userToken struct {
	scope               string
	openId              string
	unionId             string
	userAccessDeadLine  time.Time
	userRefreshDeadLine time.Time
	accessToken         string
	refreshToken        string
}

type userTokenTemp struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	UnionId      string `json:"unionid"`
	Scope        string `json:"scope"`
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}

func (ut *userToken) isAccessDead() bool {
	return time.Now().After(ut.userAccessDeadLine)
}

func (ut *userToken) isRefreshDead() bool {
	return time.Now().After(ut.userRefreshDeadLine)
}

func (ut *userToken) refresh(appid string) error {
	if ut.isAccessDead() {
		if ut.isRefreshDead() {
			return errors.New("refresh token dead")
		} else {
			var tt userTokenTemp
			param := req.Param{
				"appid":         appid,
				"grant_type":    "refresh_token",
				"refresh_token": ut.refreshToken,
			}
			err := util.RequestGetJSONWithQueryParam(refreshUserTokenURL, param, &tt)
			if err != nil {
				return err
			}
			if tt.ErrCode != 0 {
				return errors.New(strconv.Itoa(tt.ErrCode) + ":" + tt.ErrMsg)
			}
			ut.accessToken = tt.AccessToken
			ut.userAccessDeadLine = time.Now().Add(time.Second * time.Duration(tt.ExpiresIn-60))
			ut.refreshToken = tt.RefreshToken
			ut.scope = tt.Scope
			ut.openId = tt.OpenId
		}
	} else {
		return nil
	}
	return nil
}
