package util

import (
	"github.com/imroc/req"
)

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
