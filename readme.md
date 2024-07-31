# 一个用GoLang编写的微信公众号对接SDK

## 已实现功能

### 微信公众号网页授权

> https://developers.weixin.qq.com/doc/offiaccount/OA_Web_Apps/Wechat_webpage_authorization.html

可以通过CODE获取用户的基本信息，如果scope为snsapi_userinfo则可以获取用户的详细信息，如果scope为snsapi_base则只能获取用户的OpenID。

示例代码：

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

详细信息包括：

1. OpenID  OpenId
2. 昵称  Nickname
3. 性别  Sex
4. 省份  Province
5. 城市  City
6. 国家  Country
7. 头像  HeadImg
8. 特权信息  Privilege
9. UnionID  UnionId
