package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type RequestHandler struct {
}

func makeQueryStr(params map[string]string) string {
	queryStr := ""
	for k, v := range params {
		queryStr = fmt.Sprintf("%s&%s=%s", queryStr, k, v)
	}
	return queryStr
}

func (r *RequestHandler) GetRequest(url string, params map[string]string) []byte {
	queryStr := makeQueryStr(params)
	url = fmt.Sprintf("%s?%s", url, queryStr)
	resp, err := http.Get(url)
	utils.ErrorHandler(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}
