package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

func (eh *RequestHandler) EncodeToJson(data map[string]interface{}) string {

	m, _ := json.Marshal(data)
	return string(m)
}

func (eh *RequestHandler) WriteFile(data map[string]interface{}) {
	dFile, _ := os.Create("C:\\tests\\" + strconv.FormatInt(time.Now().Unix(), 10) + ".json")
	dataEcd := json.NewEncoder(dFile)
	dataEcd.Encode(data)
	dFile.Close()
}
