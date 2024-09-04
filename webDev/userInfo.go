package webDev

// UserInfo
//
// English:
//
// UserInfo represents the user information structure.
//
// It contains various fields that store user details such as OpenId, Nickname, Sex, etc.
//
// The corresponding fields are all from the user information returned by the WeChat server.
//
// Chinese:
//
// UserInfo 表示用户信息结构体。
//
// 它包含了各种字段，用于存储用户的详细信息，如 OpenId、Nickname、Sex 等。
//
// 对应的字段都是来自于微信服务器返回的用户信息。
type UserInfo struct {

	// OpenId is the unique identifier of the user.
	//
	// OpenId 是用户的唯一标识符。
	OpenId string `json:"openid"`

	// Nickname is the nickname of the user.
	//
	// Nickname 是用户的昵称。
	Nickname string `json:"nickname"`

	// Sex is the user's gender, with a value of 1 indicating male, a value of 2 indicating female, and a value of 0 indicating unknown.
	//
	// Sex 是用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Sex int `json:"sex"`

	// Province is the province where users fill in their personal information.
	//
	// Province 是用户个人资料填写的省份。
	Province string `json:"province"`

	// City is the city where ordinary users fill in their personal information.
	//
	// City 是普通用户个人资料填写的城市。
	City string `json:"city"`

	// Country is the country filled in by the user, such as CN for China.
	//
	// Country 是用户填的国家，如中国为CN。
	Country string `json:"country"`

	// HeadImg is the user's avatar, and the last value represents the size of the square avatar (there are 0, 46, 64, 96, 132 values to choose from, 0 represents 640640 square avatar). When the user does not have an avatar, this item is empty. If the user changes their avatar, the original avatar URL will become invalid.
	//
	// HeadImg 是用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	HeadImg string `json:"headimgurl"`

	// Privilege is user privilege information, JSON array, for example, WeChat Woka user is (Chinaunicom)
	//
	// Privilege 是用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	Privilege []string `json:"privilege"`

	// UnionId is the union ID of the user, which is the same as the OpenId of the same user in different WeChat applications.This field will appear only after the user binds the official account to the WeChat open platform account.
	//
	// UnionId 是用户的统一标识符，同一个用户在不同微信应用中的 OpenId 是相同的。只有在用户将公众号绑定到微信开放平台账号后，才会出现该字段。
	UnionId string `json:"unionid"`

	// ErrCode is the error code returned by the WeChat server.
	//
	// ErrCode 是微信服务器返回的错误码。
	ErrCode int `json:"errcode"`

	// ErrMsg is the error message returned by the WeChat server.
	//
	// ErrMsg 是微信服务器返回的错误消息。
	ErrMsg string `json:"errmsg"`
}
