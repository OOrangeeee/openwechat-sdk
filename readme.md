# WeChat official account docking SDK based on GoLang

> 基于GoLang的微信公众号对接SDK

This repository will be updated periodically according to the author's needs, and of course, we welcome everyone to write issues to raise their requirements. At the same time, we welcome everyone to point out shortcomings and bugs.

> 此仓库会根据作者的需求不定期更新，当然也欢迎大家写issue提出需求，同时欢迎大家指出不足和bug。


## Implemented functions

> 已实现的功能

### WeChat official account page authorization

> 微信公众号网页授权

https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html

Basic user information can be obtained through CODE. If the scope is snsapi_usinfo, detailed user information can be obtained. If the scope is snsapi_base, only the user's OpenID can be obtained.

> 可以通过CODE获取用户的基本信息，如果scope为snsapi_userinfo则可以获取用户的详细信息，如果scope为snsapi_base则只能获取用户的OpenID。

```GoLang
wt := webDev.New("appid", "appsecret", "scope")
userInfo,isDetail,err := wt.GetUserInfoByCode("code")
if err != nil {
    fmt.Println(err)
    return
}
if isDetail {
	// 可以获取用户的详细信息
    fmt.Println(userInfo)
} else {
	// 只能获取OpenID
    fmt.Println(userInfo.OpenID)
}
```

Detailed information includes:

> 详细信息包括：

1. OpenID  OpenId
2. 昵称  Nickname
3. 性别  Sex
4. 省份  Province
5. 城市  City
6. 国家  Country
7. 头像  HeadImg
8. 特权信息  Privilege
9. UnionID  UnionId

### Get automatically updated AccessToken

> 获取自动更新的AccessToken

https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html

```GoLang
wt := webDev("appid", "appsecret", "scope")
accessToken,err := wt.GetAccessToken()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(accessToken)
```

### Get automatically updated JsApiTicket

> 获取自动更新的JsApiTicket

https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/JS-SDK.html

```GoLang
wt := webDev
jsApiTicket,err := wt.GetJsApiTicket()
if err != nil {
    fmt.Println(err)
    return
}
fmt.Println(jsApiTicket)
```

