package webDev

import (
	"errors"
	"github.com/imroc/req"
	"strconv"
	"time"
)

var (
	refreshtokenURL   = "https://api.weixin.qq.com/sns/oauth2/refresh_token"
	getAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/access_token"
	getUserInfoURL    = "https://api.weixin.qq.com/sns/userinfo"
)

type WcTool struct {
	appId      string
	secret     string
	canGetInfo bool
	//users      map[string]userToken
}

func New(appId, secret, scope string) (*WcTool, error) {
	if scope == "snsapi_userinfo" {
		return &WcTool{
			appId:      appId,
			secret:     secret,
			canGetInfo: true,
			//users:      make(map[string]userToken),
		}, nil
	} else if scope == "snsapi_base" {
		return &WcTool{
			appId:      appId,
			secret:     secret,
			canGetInfo: false,
			//users:      make(map[string]userToken),
		}, nil
	} else {
		return nil, errors.New("invalid scope")
	}
}

/*func (wt *WcTool) GetUserInfoByOpenId(openId string) (*UserInfo, error) {
	if wt.canGetInfo {
		if ut, ok := wt.users[openId]; ok {
			err := ut.refresh(wt.appId)
			if err != nil {
				return nil, err
			}
			return ut.getUserInfo()
		} else {
			return nil, errors.New("no such user")
		}
	} else {
		if ut, ok := wt.users[openId]; ok {
			return &UserInfo{
				OpenId: ut.openId,
			}, nil
		} else {
			return nil, errors.New("no such user")
		}
	}
}*/

func (wt *WcTool) GetUserInfoByCode(code string) (*UserInfo, bool, error) {
	param := req.Param{
		"appid":      wt.appId,
		"secret":     wt.secret,
		"code":       code,
		"grant_type": "authorization_code",
	}
	r, err := req.Get(getAccessTokenURL, param)
	if err != nil {
		return nil, false, err
	}
	var tt tokenTemp
	err = r.ToJSON(&tt)
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
	//wt.users[ut.openId] = ut
	return ut.getUserInfo()
}
