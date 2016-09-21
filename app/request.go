package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type RequestHandler struct{}

//get request, returns body as json
func (rq *RequestHandler) DoGetRequest(url string) (m map[string]interface{}, err error) {

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	m = make(map[string]interface{})
	err = json.Unmarshal(body, &m)

	return m, err
}

//post request, returns body as json
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

//raw post request, returns body as string
func (rq *RequestHandler) DoRawPostRequest(url string, bodyType string, payload string) (b string, err error) {
	rdr := strings.NewReader(payload)

	resp, err := http.Post(url, bodyType, rdr)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	return string(body), err
}

//reads 'source' url from .env file and appends the given path
func (rq *RequestHandler) CS(url string) string {
	return os.Getenv("ELASTICSEARCH_SOURCE_URL") + url
}

//reads 'target' url from .env file and appends the given path
func (rq *RequestHandler) CT(url string) string {
	return os.Getenv("ELASTICSEARCH_TARGET_URL") + url
}

func (eh *RequestHandler) EncodeToJson(data map[string]interface{}) string {

	m, _ := json.Marshal(data)
	return string(m)
}

//saves given map data to file
func (eh *RequestHandler) WriteFile(fullname string, data map[string]interface{}) {
	dFile, _ := os.Create(fullname)
	dataEcd := json.NewEncoder(dFile)
	dataEcd.Encode(data)
	dFile.Close()
}
