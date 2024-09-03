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
	accessTokenURL        = "https://api.weixin.qq.com/cgi-bin/token"
	jsapiTicketURL        = "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
)

type WcTool struct {
	appId       string
	secret      string
	canGetInfo  bool
	accessToken *accessToken
	jsapiTicket *jsapiTicket
}

func New(appId, secret, scope string) (*WcTool, error) {
	var canGetInfo bool
	if scope == "snsapi_userinfo" {
		canGetInfo = true
	} else if scope == "snsapi_base" {
		canGetInfo = false
	} else {
		return nil, errors.New("invalid scope")
	}
	return &WcTool{
		appId:      appId,
		secret:     secret,
		canGetInfo: canGetInfo,
		accessToken: &accessToken{
			token:    "",
			deadLine: time.Now(),
		},
		jsapiTicket: &jsapiTicket{
			ticket:   "",
			deadLine: time.Now(),
		},
	}, nil
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

func (wt *WcTool) GetAccessToken() (string, error) {
	return wt.accessToken.getAccessToken(wt.appId, wt.secret)
}

func (wt *WcTool) GetJsapiTicket() (string, error) {
	token, err := wt.GetAccessToken()
	if err != nil {
		return "", err
	}
	return wt.jsapiTicket.getJsapiTicket(token)
}

func (wt *WcTool) GetUserInfoByCode(code string) (*UserInfo, bool, error) {
	var tt userTokenTemp
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
		accessToken:         tt.AccessToken,
		userAccessDeadLine:  time.Now().Add(time.Second * time.Duration(tt.ExpiresIn-60)),
		refreshToken:        tt.RefreshToken,
		userRefreshDeadLine: time.Now().Add(30 * 24 * time.Hour),
		scope:               tt.Scope,
		openId:              tt.OpenId,
		unionId:             tt.UnionId,
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
