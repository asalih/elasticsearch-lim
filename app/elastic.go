package app

import (
	"fmt"
)

type ElasticHandler struct{ LastReq *RequestResult }

type RequestResult struct{ Result map[string]interface{} }

func (eh *ElasticHandler) CollectNewData() {

	req := &RequestHandler{}

	mainGet := req.DoGetRequest(req.CS("/_stats"))
	result := &RequestResult{Result: mainGet}

	eh.ProcessNewData(result)
	eh.LastReq = result

	//fmt.Println(eh.cnt)
}

func (eh *ElasticHandler) ProcessNewData(result *RequestResult) {
	//	fmt.Println(result.Result["_all"])
	if eh.LastReq == nil {
		eh.LastReq = result
	}
	lastData := eh.LastReq.Result["_all"].(map[string]interface{})["total"]
	data := result.Result["_all"].(map[string]interface{})["total"]

	eh.DoCalculations(data.(map[string]interface{}), lastData.(map[string]interface{}), "_all")

	fmt.Println(data)
}

func (eh *ElasticHandler) DoCalculations(data map[string]interface{}, lastData map[string]interface{}, idx string) {

	object := make(map[string]interface{})
	object["idx"] = idx

}
