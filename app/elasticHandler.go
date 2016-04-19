package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type ElasticHandler struct{}

func (eh *ElasticHandler) CollectNewData() {

	mainGet := eh.DoGetRequest(eh.CS("/"))
	fmt.Println(mainGet["cluster_name"])

	subs := mainGet["version"]
	sub := subs.(map[string]interface{})
	fmt.Println(sub["number"])
}

func (eh *ElasticHandler) DoGetRequest(url string) map[string]interface{} {

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

func (eh *ElasticHandler) DoPostRequest(url string, bodyType string, payload string) map[string]interface{} {
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

func (eh *ElasticHandler) CS(url string) string {
	return os.Getenv("ELASTICSEARCH_SOURCE_URL") + url
}

func (eh *ElasticHandler) CT(url string) string {
	return os.Getenv("ELASTICSEARCH_TARGET_URL") + url
}
