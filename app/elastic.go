package app

import (
	"os"
	"strconv"
	"time"
)

type ElasticHandler struct {
	LastReq *RequestResult
	Time    time.Time
}

var nonCalculationFields = []string{""}

type RequestResult struct{ Result map[string]interface{} }

func (eh *ElasticHandler) CollectNewData() {

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
}

func (eh *ElasticHandler) DoCalculations(data map[string]interface{}, lastData map[string]interface{}, idx string) map[string]interface{} {

	object := make(map[string]interface{})
	object["idx"] = idx
	object["timestamp"] = eh.Time.Unix()

	for i, _ := range data {
		hNew := data[i].(map[string]interface{})
		hOld := lastData[i].(map[string]interface{})

		for j, _ := range hNew {
			fl, ok := hNew[j].(float64)
			if ok {
				flo, _ := hOld[j].(float64)

				object[i+"."+j] = eh.Calc(fl, flo, j)

			}
		}
	}

	//eh.WriteFile(object)

	//fmt.Println(object)

	return object
}

func (eh *ElasticHandler) Calc(f1 float64, f2 float64, field string) float64 {
	switch field {
	case "count", "deleted", "size_in_bytes", "delete_total", "total", "open_contexts":
		return f1
	}
	intv, _ := strconv.ParseFloat(os.Getenv("INTERVAL_SECOND"), 10)
	return (f1 - f2) / intv
}
