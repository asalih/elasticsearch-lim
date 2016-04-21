package app

import (
	"fmt"
	"os"
	"time"
)

type ElasticHandler struct{ LastReq *RequestResult }

type RequestResult struct{ Result map[string]interface{} }

func (eh *ElasticHandler) CollectNewData(time time.Time) {

	req := &RequestHandler{}

	mainGet := req.DoGetRequest(req.CS("/_stats"))
	result := &RequestResult{Result: mainGet}

	eh.ProcessNewData(result, req)
	eh.LastReq = result

	//fmt.Println(eh.cnt)
}

func (eh *ElasticHandler) ProcessNewData(result *RequestResult, reqHandler *RequestHandler) {

	if eh.LastReq == nil {
		eh.LastReq = result
	}
	targetUrl := reqHandler.CT(os.Getenv("ELASTICSEARCH_INDEX"))

	lastData := eh.LastReq.Result["_all"].(map[string]interface{})["total"].(map[string]interface{})
	data := result.Result["_all"].(map[string]interface{})["total"].(map[string]interface{})

	all := eh.DoCalculations(data, lastData, "_all")
	reqHandler.DoPostRequest(targetUrl, "text/json", reqHandler.EncodeToJson(all))

	forLastData := eh.LastReq.Result["indices"].(map[string]interface{})
	forData := result.Result["indices"].(map[string]interface{})

	for i, _ := range forData {
		ld := forLastData[i]
		if ld != nil {
			lastData = ld.(map[string]interface{})["total"].(map[string]interface{})
			data = forData[i].(map[string]interface{})["total"].(map[string]interface{})

			indices := eh.DoCalculations(data, lastData, i)
			reqHandler.DoPostRequest(targetUrl, "text/json", reqHandler.EncodeToJson(indices))
		}
	}
	fmt.Println("tick")
}

func (eh *ElasticHandler) DoCalculations(data map[string]interface{}, lastData map[string]interface{}, idx string) map[string]interface{} {

	object := make(map[string]interface{})
	object["idx"] = idx
	object["timestamp"] = time.Now().Unix()

	for i, _ := range data {
		hNew := data[i].(map[string]interface{})
		hOld := lastData[i].(map[string]interface{})

		for j, _ := range hNew {
			fl, ok := eh.ToFloat(hNew[j])
			if ok {
				flo, _ := eh.ToFloat(hOld[j])
				object[i+"."+j] = eh.Calc(fl, flo)
			}
		}
	}

	//eh.WriteFile(object)

	//fmt.Println(object)

	return object
}

func (eh *ElasticHandler) Calc(f1 float64, f2 float64) float64 {
	return (f1 - f2) / 120
}

func (eh *ElasticHandler) ToFloat(data interface{}) (float64, bool) {
	f, o := data.(float64)
	return f, o
}
