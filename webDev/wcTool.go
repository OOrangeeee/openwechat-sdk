package webDev

import (
	"errors"
	"github.com/OOrangeeee/openwechat-sdk/webDev/util"
	"github.com/imroc/req"
	"strconv"
	"time"
)

var (
	refreshUserTokenURL   = "https://api.weixin.qq.com/sns/oauth2/refresh_token"
	getUserAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/access_token"
	getUserInfoURL        = "https://api.weixin.qq.com/sns/userinfo"
)

type WcTool struct {
	appId      string
	secret     string
	canGetInfo bool
}

func New(appId, secret, scope string) (*WcTool, error) {
	if scope == "snsapi_userinfo" {
		return &WcTool{
			appId:      appId,
			secret:     secret,
			canGetInfo: true,
		}, nil
	} else if scope == "snsapi_base" {
		return &WcTool{
			appId:      appId,
			secret:     secret,
			canGetInfo: false,
		}, nil
	} else {
		return nil, errors.New("invalid scope")
	}
}

func (wt *WcTool) SetAppId(newAppId string) {
	wt.appId = newAppId
}

func (wt *WcTool) SetSecret(newSecret string) {
	wt.secret = newSecret
}

func (wt *WcTool) SetCanGetInfo(newScope string) {
	if newScope == "snsapi_userinfo" {
		wt.canGetInfo = true
	} else if newScope == "snsapi_base" {
		wt.canGetInfo = false
	}
}

func (wt *WcTool) GetUserInfoByCode(code string) (*UserInfo, bool, error) {
	var tt tokenTemp
	param := req.Param{
		"appid":      wt.appId,
		"secret":     wt.secret,
		"code":       code,
		"grant_type": "authorization_code",
	}
	err := util.RequestGetJSONWithQueryParam(getUserAccessTokenURL, param, &tt)
	if err != nil {
		return nil, false, err
	}
	if tt.ErrCode != 0 {
		return nil, false, errors.New(strconv.Itoa(tt.ErrCode) + ":" + tt.ErrMsg)
	}
	ut := userToken{
		accessToken:     tt.AccessToken,
		accessDeadLine:  time.Now().Add(time.Second * time.Duration(tt.ExpiresIn)),
		refreshToken:    tt.RefreshToken,
		refreshDeadLine: time.Now().Add(30 * 24 * time.Hour),
		scope:           tt.Scope,
		openId:          tt.OpenId,
		unionId:         tt.UnionId,
	}
	return wt.getUserInfo(&ut)
}

func (wt *WcTool) getUserInfo(ut *userToken) (*UserInfo, bool, error) {
	if !wt.canGetInfo {
		return &UserInfo{
			OpenId: ut.openId,
		}, false, nil
	}
	var ui UserInfo
	param := req.Param{
		"access_token": ut.accessToken,
		"openid":       ut.openId,
		"lang":         "zh_CN",
	}
	err := util.RequestGetJSONWithQueryParam(getUserInfoURL, param, &ui)
	if err != nil {
		return nil, false, err
	}
	if ui.ErrCode != 0 {
		return nil, false, errors.New(strconv.Itoa(ui.ErrCode) + ":" + ui.ErrMsg)
	}
	return &ui, true, nil
}
