package util

import (
	"github.com/imroc/req"
)

// RequestGetJSONWithQueryParam
//
// English:
//
// RequestGetJSONWithQueryParam is a function that sends a GET request to the specified URL with the specified parameters and returns the result in JSON format.
//
// url: The URL to send the request to.
//
// param: The parameters to send with the request.
//
// v: The variable to store the result.
//
// Returns an error if the request fails.
//
// Chinese:
//
// RequestGetJSONWithQueryParam 是一个发送 GET 请求到指定 URL 的函数，带有指定的参数，并以 JSON 格式返回结果。
//
// url: 发送请求的 URL。
//
// param: 发送请求的参数。
//
// v: 存储结果的变量。
//
// 如果请求失败，返回一个错误。
func RequestGetJSONWithQueryParam(url string, param req.Param, v interface{}) error {
	r, err := req.Get(url, param)
	if err != nil {
		return err
	}
	err = r.ToJSON(v)
	if err != nil {
		return err
	}
	return nil
}
