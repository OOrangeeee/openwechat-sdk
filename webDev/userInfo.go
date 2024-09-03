package webDev

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
