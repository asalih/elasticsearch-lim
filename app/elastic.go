package app

import (
	"os"
	"strconv"
	"time"

	"github.com/asalih/elasticsearch-lim/q"
)

type ElasticHandler struct {
	LastReq *RequestResult
	Time    time.Time
}

var nonCalculationFields = []string{""}

type RequestResult struct {
	StatsResult    map[string]interface{}
	StatsTargetUrl string
	NodesResult    map[string]interface{}
	NodesTargetUrl string
}

func (eh *ElasticHandler) CheckMappings() {
	req := &RequestHandler{}

	mapUrl := req.CT(os.Getenv("ELASTICSEARCH_INDEX_ST") + "_mapping")
	result := req.DoGetRequest(mapUrl)
	status := result["status"]

	if len(result) == 0 || (status != nil && status.(float64) == 404) {

		qTemp := q.QueryTemplates{}

		template := qTemp.GetTargetMappingTemplate(os.Getenv("ELASTICSEARCH_TTL"))
		req.DoPostRequest(mapUrl, "text/json", template)
	}

	mapUrl = req.CT(os.Getenv("ELASTICSEARCH_INDEX_ND") + "_mapping")
	result = req.DoGetRequest(mapUrl)
	status = result["status"]

	if len(result) == 0 || (status != nil && status.(float64) == 404) {

		qTemp := q.QueryTemplates{}

		template := qTemp.GetTargetMappingTemplate(os.Getenv("ELASTICSEARCH_TTL"))
		req.DoPostRequest(mapUrl, "text/json", template)
	}
}

func (eh *ElasticHandler) CollectNewData() {

	req := &RequestHandler{}

	statsGet := req.DoGetRequest(req.CS("/_stats"))
	result := &RequestResult{StatsResult: statsGet, StatsTargetUrl: req.CT(os.Getenv("ELASTICSEARCH_INDEX_ST"))}

	eh.ProcessStatsData(result, req)

	nodesGet := req.DoGetRequest(req.CS("/_nodes/stats"))
	result.NodesResult = nodesGet
	result.NodesTargetUrl = req.CT(os.Getenv("ELASTICSEARCH_INDEX_ND"))

	eh.ProcessNodesData(result, req)

	eh.LastReq = result

}

func (eh *ElasticHandler) ProcessStatsData(result *RequestResult, reqHandler *RequestHandler) {

	if eh.LastReq == nil {
		eh.LastReq = result
	}
	//targetUrl := reqHandler.CT(os.Getenv("ELASTICSEARCH_INDEX_ST"))

	lastData := eh.LastReq.StatsResult["_all"].(map[string]interface{})["total"].(map[string]interface{})
	data := result.StatsResult["_all"].(map[string]interface{})["total"].(map[string]interface{})

	all := make(map[string]interface{})
	eh.DoCalculations(data, lastData, "_all", all, "")

	reqHandler.DoPostRequest(result.StatsTargetUrl, "text/json", reqHandler.EncodeToJson(all))

	forLastData := eh.LastReq.StatsResult["indices"].(map[string]interface{})
	forData := result.StatsResult["indices"].(map[string]interface{})

	for i, _ := range forData {
		ld := forLastData[i]
		if ld != nil {
			lastData = ld.(map[string]interface{})["total"].(map[string]interface{})
			data = forData[i].(map[string]interface{})["total"].(map[string]interface{})

			indices := make(map[string]interface{})
			eh.DoCalculations(data, lastData, i, indices, "")
			reqHandler.DoPostRequest(result.StatsTargetUrl, "text/json", reqHandler.EncodeToJson(indices))
		}
	}
}

func (eh *ElasticHandler) ProcessNodesData(result *RequestResult, reqHandler *RequestHandler) {

	if eh.LastReq == nil {
		eh.LastReq = result
	}
	//targetUrl := reqHandler.CT(os.Getenv("ELASTICSEARCH_INDEX_ST"))
	forLastData := eh.LastReq.NodesResult["nodes"].(map[string]interface{})
	forData := result.NodesResult["nodes"].(map[string]interface{})

	for i, _ := range forData {
		ld := forLastData[i]
		if ld != nil {

			fCurr := forData[i].(map[string]interface{})

			jvmLastData := ld.(map[string]interface{})["jvm"].(map[string]interface{})
			jvmData := fCurr["jvm"].(map[string]interface{})

			osLastData := ld.(map[string]interface{})["os"].(map[string]interface{})
			osData := fCurr["os"].(map[string]interface{})

			httpLastData := ld.(map[string]interface{})["http"].(map[string]interface{})
			httpData := fCurr["http"].(map[string]interface{})

			nodes := make(map[string]interface{})

			name := fCurr["name"].(string)
			eh.DoCalculations(jvmData, jvmLastData, name, nodes, "jvm")
			eh.DoCalculations(osData, osLastData, name, nodes, "os")
			eh.DoCalculations(httpData, httpLastData, name, nodes, "http")

			nodes["node_id"] = i
			nodes["transport_address"] = fCurr["transport_address"].(string)

			reqHandler.DoPostRequest(result.NodesTargetUrl, "text/json", reqHandler.EncodeToJson(nodes))

		}
	}
}

func (eh *ElasticHandler) DoCalculations(data map[string]interface{}, lastData map[string]interface{}, idx string, object map[string]interface{}, initial string) {

	object["idx"] = idx
	object["timestamp"] = eh.Time.Unix()

	for i, _ := range data {
		hNew, ok := data[i].(map[string]interface{})
		hOld, _ := lastData[i].(map[string]interface{})
		if ok {
			for j, _ := range hNew {
				fl, ok := hNew[j].(float64)
				if ok {
					flo, _ := hOld[j].(float64)
					field := ""
					if initial != "" {
						field = initial + "." + i + "." + j
					} else {
						field = i + "." + j
					}
					object[field] = eh.Calc(fl, flo, field)

				}
			}
		} else {
			field := ""
			if initial != "" {
				field = initial + "." + i
			} else {
				field = i
			}
			object[field] = eh.Calc(data[i].(float64), lastData[i].(float64), field)
		}

	}

}

func (eh *ElasticHandler) Calc(f1 float64, f2 float64, field string) float64 {
	switch field {
	case "docs.count", "docs.deleted", "store.size_in_bytes", "indexing.delete_total", "search.open_contexts", "http.current_open", "os.timestamp",
		"jvm.mem.heap_used_in_bytes", "jvm.mem.heap_used_percent", "os.cpu_percent", "os.load_average", "os.mem.total_in_bytes", "os.mem.used_percent",
		"os.mem.free_percent":
		return f1
	}
	intv, _ := strconv.ParseFloat(os.Getenv("INTERVAL_SECOND"), 10)
	return (f1 - f2) / intv
}
