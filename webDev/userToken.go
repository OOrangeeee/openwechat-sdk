package webDev

import (
	"errors"
	"github.com/imroc/req"
	"strconv"
	"time"
)

type userToken struct {
	scope           string
	openId          string
	unionId         string
	accessDeadLine  time.Time
	refreshDeadLine time.Time
	accessToken     string
	refreshToken    string
}

type tokenTemp struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	UnionId      string `json:"unionid"`
	Scope        string `json:"scope"`
	ErrCode      int    `json:"errcode"`
	ErrMsg       string `json:"errmsg"`
}

type UserInfo struct {
	OpenId    string   `json:"openid"`
	Nickname  string   `json:"nickname"`
	Sex       int      `json:"sex"`
	Province  string   `json:"province"`
	City      string   `json:"city"`
	Country   string   `json:"country"`
	HeadImg   string   `json:"headimgurl"`
	Privilege []string `json:"privilege"`
	UnionId   string   `json:"unionid"`
	ErrCode   int      `json:"errcode"`
	ErrMsg    string   `json:"errmsg"`
}

func (ut *userToken) isAccessDead() bool {
	return time.Now().After(ut.accessDeadLine)
}

func (ut *userToken) isRefreshDead() bool {
	return time.Now().After(ut.refreshDeadLine)
}

func (ut *userToken) refresh(appid string) error {
	if ut.isAccessDead() {
		if ut.isRefreshDead() {
			return errors.New("refresh token dead")
		} else {
			param := req.Param{
				"appid":         appid,
				"grant_type":    "refresh_token",
				"refresh_token": ut.refreshToken,
			}
			r, err := req.Get(refreshtokenURL, param)
			if err != nil {
				return err
			}
			var tt tokenTemp
			err = r.ToJSON(&tt)
			if err != nil {
				return err
			}
			if tt.ErrCode != 0 {
				return errors.New(strconv.Itoa(tt.ErrCode) + ":" + tt.ErrMsg)
			}
			ut.accessToken = tt.AccessToken
			ut.accessDeadLine = time.Now().Add(time.Second * time.Duration(tt.ExpiresIn))
			ut.refreshToken = tt.RefreshToken
			ut.scope = tt.Scope
			ut.openId = tt.OpenId
		}
	} else {
		return nil
	}
	return nil
}

func (ut *userToken) getUserInfo() (*UserInfo, bool, error) {
	if ut.scope != "snsapi_userinfo" {
		return &UserInfo{
			OpenId: ut.openId,
		}, false, nil
	}
	param := req.Param{
		"access_token": ut.accessToken,
		"openid":       ut.openId,
		"lang":         "zh_CN",
	}
	r, err := req.Get(getUserInfoURL, param)
	if err != nil {
		return nil, false, err
	}
	var ui UserInfo
	err = r.ToJSON(&ui)
	if err != nil {
		return nil, false, err
	}
	if ui.ErrCode != 0 {
		return nil, false, errors.New(strconv.Itoa(ui.ErrCode) + ":" + ui.ErrMsg)
	}
	return &ui, true, nil
}
