package wechat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Post send post request.
func Post(url, contentType string, args map[string]string) (result string, err error) {
	if contentType == "" {
		contentType = "application/x-www-form-urlencoded"
	}
	resp, err := http.Post(url, contentType, strings.NewReader(argsEncode(args)))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = NewError(body)
	if err != nil {
		return
	}
	result = string(body)
	return
}

// PostJSON send post request.
func PostJSON(url string, jsonObject interface{}) (result string, err error) {
	json, err := json.Marshal(jsonObject)
	if err != nil {
		return
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(json))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = NewError(body)
	if err != nil {
		return
	}
	result = string(body)
	return
}

// Get send get request.
func Get(url string, args map[string]string) (result []byte, err error) {
	if args != nil {
		url += "?" + argsEncode(args)
	}
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = NewError(body)
	if err != nil {
		return
	}
	result = body
	return
}

func argsEncode(params map[string]string) string {
	args := url.Values{}
	for k, v := range params {
		args.Set(k, v)
	}
	return args.Encode()
}
