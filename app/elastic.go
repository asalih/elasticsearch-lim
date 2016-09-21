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

type RequestResult struct {
	StatsResult    map[string]interface{}
	StatsTargetUrl string
	NodesResult    map[string]interface{}
	NodesTargetUrl string
}

//mapping checker, if mapping doesn't exists, posts new one for 'ttl' properties.
func (eh *ElasticHandler) CheckMappings() {
	req := &RequestHandler{}

	mapUrl := req.CT(os.Getenv("ELASTICSEARCH_INDEX_ST") + "_mapping")
	result, err := req.DoGetRequest(mapUrl)

	if err != nil {
		return
	}

	status := result["status"]

	if len(result) == 0 || (status != nil && status.(float64) == 404) {

		qTemp := q.QueryTemplates{}

		template := qTemp.GetTargetMappingTemplate(os.Getenv("ELASTICSEARCH_TTL"))
		req.DoPostRequest(mapUrl, "text/json", template)
	}

	mapUrl = req.CT(os.Getenv("ELASTICSEARCH_INDEX_ND") + "_mapping")
	result, err = req.DoGetRequest(mapUrl)
	status = result["status"]

	if len(result) == 0 || (status != nil && status.(float64) == 404) {

		qTemp := q.QueryTemplates{}

		template := qTemp.GetTargetMappingTemplate(os.Getenv("ELASTICSEARCH_TTL"))
		req.DoPostRequest(mapUrl, "text/json", template)
	}
}

//collects index stats and node stats.
func (eh *ElasticHandler) CollectNewData() {

	req := &RequestHandler{}

	statsGet, err := req.DoGetRequest(req.CS("/_stats"))
	if err != nil {
		return
	}
	result := &RequestResult{StatsResult: statsGet, StatsTargetUrl: req.CT(os.Getenv("ELASTICSEARCH_INDEX_ST"))}

	eh.ProcessStatsData(result, req)

	nodesGet, err := req.DoGetRequest(req.CS("/_nodes/stats"))
	result.NodesResult = nodesGet
	result.NodesTargetUrl = req.CT(os.Getenv("ELASTICSEARCH_INDEX_ND"))

	eh.ProcessNodesData(result, req)

	eh.LastReq = result

}

//processes collected new index stats data
func (eh *ElasticHandler) ProcessStatsData(result *RequestResult, reqHandler *RequestHandler) {

	if eh.LastReq == nil {
		eh.LastReq = result
	}

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

//processes collected new nodes stats data
func (eh *ElasticHandler) ProcessNodesData(result *RequestResult, reqHandler *RequestHandler) {

	if eh.LastReq == nil {
		eh.LastReq = result
	}

	forLastData := eh.LastReq.NodesResult["nodes"].(map[string]interface{})
	forData := result.NodesResult["nodes"].(map[string]interface{})

	for i, _ := range forData {
		ld := forLastData[i]
		if ld != nil {

			fCurr := forData[i].(map[string]interface{})

			nodes := make(map[string]interface{})

			name := fCurr["name"].(string)
			eh.DoCalculations(fCurr, ld.(map[string]interface{}), name, nodes, "")

			reqHandler.DoPostRequest(result.NodesTargetUrl, "text/json", reqHandler.EncodeToJson(nodes))

		}
	}
}

//Data is the last http result, lastData is for comparison
func (eh *ElasticHandler) DoCalculations(data map[string]interface{}, lastData map[string]interface{}, idx string, object map[string]interface{}, initial string) {

	object["idx"] = idx
	object["timestamp"] = eh.Time.Unix()

	for i, _ := range data {

		hNew, ok := data[i].(map[string]interface{})
		hOld, _ := lastData[i].(map[string]interface{})
		if ok {
			for j, _ := range hNew {
				field := ""
				if initial != "" {
					field = initial + "." + i + "." + j
				} else {
					field = i + "." + j
				}

				fl, ok := hNew[j].(float64)
				if ok {
					flo, _ := hOld[j].(float64)

					object[field] = eh.Calc(fl, flo, field)
				} else {
					sub, oks := hNew[j].(map[string]interface{})
					if oks {

						eh.DoCalculations(sub, hOld[j].(map[string]interface{}), idx, object, field)
					}
				}
			}
		} else {
			field := ""
			if initial != "" {
				field = initial + "." + i
			} else {
				field = i
			}
			df1, isFl := data[i].(float64)
			if isFl {
				object[field] = eh.Calc(df1, lastData[i].(float64), field)
			} else {
				ds1, isStr := data[i].(string)
				if isStr {
					object[field] = ds1
				}
			}
		}
	}
}

//takes the diffrence given two(f1,f2) values and divides interval seconds(lookup the .env file). accept some fields, lookup switch-case closure.
func (eh *ElasticHandler) Calc(f1 float64, f2 float64, field string) float64 {
	switch field {
	case "docs.count", "docs.deleted", "store.size_in_bytes", "indexing.delete_total", "search.open_contexts", "http.current_open", "os.timestamp",
		"jvm.mem.heap_used_in_bytes", "jvm.mem.heap_used_percent", "os.cpu_percent", "os.load_average", "os.mem.total_in_bytes", "os.mem.used_percent",
		"os.mem.free_percent", "jvm.mem.non_heap_used_in_bytes", "jvm.timestamp", "gc.collectors.old.collection_count",
		"gc.collectors.young.collection_count", "os.mem.used_in_bytes", "transport.server_open":
		return f1
	}
	intv, _ := strconv.ParseFloat(os.Getenv("INTERVAL_SECOND"), 10)
	return (f1 - f2) / intv
}
