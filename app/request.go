package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type RequestHandler struct{}

func (rq *RequestHandler) DoGetRequest(url string) map[string]interface{} {

	resp, _ := http.Get(url)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	m := make(map[string]interface{})
	err := json.Unmarshal(body, &m)
	if err != nil {
		panic(err)
	}
	return m
}

func (rq *RequestHandler) DoPostRequest(url string, bodyType string, payload string) map[string]interface{} {
	rdr := strings.NewReader(payload)

	resp, _ := http.Post(url, bodyType, rdr)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	m := make(map[string]interface{})
	err := json.Unmarshal(body, &m)
	if err != nil {
		panic(err)
	}
	return m
}

func (rq *RequestHandler) CS(url string) string {
	return os.Getenv("ELASTICSEARCH_SOURCE_URL") + url
}

func (rq *RequestHandler) CT(url string) string {
	return os.Getenv("ELASTICSEARCH_TARGET_URL") + url
}
