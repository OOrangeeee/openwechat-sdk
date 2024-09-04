package webDev

import (
	"errors"
	"github.com/OOrangeeee/openwechat-sdk/webDev/util"
	"github.com/imroc/req"
	"strconv"
	"time"
)

var (

	// refreshUserTokenURL is the WeChat URL for refreshing the userAccessToken. refreshUserTokenURL 是刷新用户AccessToken的微信URL。
	refreshUserTokenURL = "https://api.weixin.qq.com/sns/oauth2/refresh_token"

	// getUserAccessTokenURL is the WeChat URL for getting the userToken. getUserAccessTokenURL 是获取用户Token的微信URL。
	getUserAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/access_token"

	// getUserInfoURL is the WeChat URL for getting user information. getUserInfoURL 是获取用户信息的微信URL。
	getUserInfoURL = "https://api.weixin.qq.com/sns/userinfo"

	// accessTokenURL is the WeChat URL for getting the accessToken. accessTokenURL 是获取accessToken的微信URL。
	accessTokenURL = "https://api.weixin.qq.com/cgi-bin/token"

	// jsapiTicketURL is the WeChat URL for getting the jsapiTicket. jsapiTicketURL 是获取jsapiTicket的微信URL
	jsapiTicketURL = "https://api.weixin.qq.com/cgi-bin/ticket/getticket"
)

// WcTool
//
// English:
//
// WcPool is the core structure of the entire package. All functions are encapsulated here.
//
// Currently supported functions:
//
// 1. Get AccessToken
//
// 2. Get JsapiTicket
//
// 3. Get user information.
//
// 4. Set and get appid, secret, and scope.
//
// Chinese:
//
// WcTool 是整个包的核心结构体。封装了所有功能在此。
//
// 目前支持的功能:
//
// 1. 获取 AccessToken
//
// 2. 获取 JsapiTicket
//
// 3. 获取用户信息。
//
// 4. 对appid、secret、scope的设置和获取。
type WcTool struct {
	appId       string
	secret      string
	canGetInfo  bool
	accessToken *accessToken
	jsapiTicket *jsapiTicket
}

// New
//
// English:
//
// New is a constructor for WcTool. It initializes the WcTool structure and returns a pointer to the structure.
//
// appId: The appid of the WeChat public account.
//
// secret: The secret of the WeChat public account.
//
// scope: The scope of the WeChat public account. It can be "snsapi_userinfo" or "snsapi_base".
//
// If the scope is "snsapi_userinfo", the user information can be obtained. If the scope is "snsapi_base", the user information cannot be obtained.
//
// Chinese:
//
// New 是 WcTool 的构造函数。它初始化 WcTool 结构体并返回结构体的指针。
//
// appId: 微信公众号的 appid。
//
// secret: 微信公众号的 secret。
//
// scope: 微信公众号的 scope。可以是 "snsapi_userinfo" 或 "snsapi_base"。
//
// 如果 scope 是 "snsapi_userinfo"，则可以获取用户详细信息。如果 scope 是 "snsapi_base"，则无法获取用户详细信息。
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

// Getter&&Setter

// GetAppId
//
// English:
//
// GetAppId is a getter method for the appId field.
//
// Chinese:
//
// GetAppId 是 appId 字段的 getter 方法。
func (wt *WcTool) GetAppId() string {
	return wt.appId
}

// GetSecret
//
// English:
//
// GetSecret is a getter method for the secret field.
//
// Chinese:
//
// GetSecret 是 secret 字段的 getter 方法。
func (wt *WcTool) GetSecret() string {
	return wt.secret
}

// GetCanGetInfo
//
// English:
//
// GetCanGetInfo is a getter method for the canGetInfo field.
//
// Chinese:
//
// GetCanGetInfo 是 canGetInfo 字段的 getter 方法。
func (wt *WcTool) GetCanGetInfo() bool {
	return wt.canGetInfo
}

// SetAppId
//
// English:
//
// SetAppId is a setter method for the appId field.
//
// Chinese:
//
// SetAppId 是 appId 字段的 setter 方法。
func (wt *WcTool) SetAppId(newAppId string) {
	wt.appId = newAppId
}

// SetSecret
//
// English:
//
// SetSecret is a setter method for the secret field.
//
// Chinese:
//
// SetSecret 是 secret 字段的 setter 方法。
func (wt *WcTool) SetSecret(newSecret string) {
	wt.secret = newSecret
}

// SetCanGetInfo
//
// English:
//
// SetCanGetInfo is a setter method for the canGetInfo field.
//
// Chinese:
//
// SetCanGetInfo 是 canGetInfo 字段的 setter 方法。
func (wt *WcTool) SetCanGetInfo(newScope string) {
	if newScope == "snsapi_userinfo" {
		wt.canGetInfo = true
	} else if newScope == "snsapi_base" {
		wt.canGetInfo = false
	}
}

// GetAccessToken
//
// English:
//
// GetAccessToken is a method that gets the AccessToken.
//
// It calls the getAccessToken method of the accessToken structure to get the AccessToken.
//
// If the AccessToken is not expired, it will return the AccessToken directly.
//
// If the AccessToken is expired, it will call the refreshAccessToken method of the accessToken structure to refresh the AccessToken.
//
// Chinese:
//
// GetAccessToken 是获取 AccessToken 的方法。
//
// 它调用 accessToken 结构体的 getAccessToken 方法来获取 AccessToken。
//
// 如果 AccessToken 没有过期，它将直接返回 AccessToken。
//
// 如果 AccessToken 过期，它将调用 accessToken 结构体的 refreshAccessToken 方法来刷新 AccessToken。
func (wt *WcTool) GetAccessToken() (string, error) {
	return wt.accessToken.getAccessToken(wt.appId, wt.secret)
}

// GetJsapiTicket
//
// English:
//
// GetJsapiTicket is a method that gets the JsapiTicket.
//
// It calls the getJsapiTicket method of the jsapiTicket structure to get the JsapiTicket.
//
// If the JsapiTicket is not expired, it will return the JsapiTicket directly.
//
// If the JsapiTicket is expired, it will call the refreshJsapiTicket method of the jsapiTicket structure to refresh the JsapiTicket.
//
// Chinese:
//
// GetJsapiTicket 是获取 JsapiTicket 的方法。
//
// 它调用 jsapiTicket 结构体的 getJsapiTicket 方法来获取 JsapiTicket。
//
// 如果 JsapiTicket 没有过期，它将直接返回 JsapiTicket。
//
// 如果 JsapiTicket 过期，它将调用 jsapiTicket 结构体的 refreshJsapiTicket 方法来刷新 JsapiTicket。
func (wt *WcTool) GetJsapiTicket() (string, error) {
	token, err := wt.GetAccessToken()
	if err != nil {
		return "", err
	}
	return wt.jsapiTicket.getJsapiTicket(token)
}

// GetUserInfoByCode
//
// English:
//
// GetUserInfoByCode is a method to get user information by code.
//
// It sends a request to get userToken.
//
// Then it calls the getUserInfo method to get user information.
//
// If the scope is "snsapi_userinfo", detailed user information can be obtained. If the scope is "snsapi_base", OpenId can be obtained.
//
// Chinese:
//
// GetUserInfoByCode 是通过 code 获取用户信息的方法。
//
// 它发送请求获取 userToken。
//
// 然后调用 getUserInfo 方法来获取用户信息。
//
// 如果 scope 是 "snsapi_userinfo"，则可以获取用户详细信息。如果 scope 是 "snsapi_base"，则获取用户OpenId。
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

// getUserInfo
//
// English:
//
// getUserInfo is a method to get user information.
//
// It sends a request to get user information.
//
// If the scope is "snsapi_userinfo", detailed user information can be obtained. If the scope is "snsapi_base", OpenId can be obtained.
//
// Chinese:
//
// getUserInfo 是获取用户信息的方法。
//
// 它发送请求获取用户信息。
//
// 如果 scope 是 "snsapi_userinfo"，则可以获取用户详细信息。如果 scope 是 "snsapi_base"，则获取用户OpenId。
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
