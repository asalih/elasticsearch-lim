package app

import (
	"os"

	"github.com/asalih/elasticsearch-lim/q"
)

type ChartData struct{}

func (c *ChartData) GetData(unixT int64, idx string, env string) string {
	qt := &q.QueryTemplates{}
	qs := qt.GetCurrentQueryTemplate(unixT, idx)

	req := &RequestHandler{}
	result, _ := req.DoRawPostRequest(req.CT(os.Getenv("ELASTICSEARCH_INDEX_"+env)+"_search"), "text/json", qs)

	return result

}

func (c *ChartData) GetLoadData(unixT int64) (json string, err error) {
	qt := &q.QueryTemplates{}
	qs := qt.GetIndicesTemplate(unixT)

	req := &RequestHandler{}
	result, errR := req.DoRawPostRequest(req.CT(os.Getenv("ELASTICSEARCH_INDEX_ST")+"_search"), "text/json", qs)
	resultN, errN := req.DoRawPostRequest(req.CT(os.Getenv("ELASTICSEARCH_INDEX_ND")+"_search"), "text/json", qs)

	if errR != nil {
		err = errR
	} else if errN != nil {
		err = errN
	}

	json = "{\"s\": " + result + ", \"n\": " + resultN + "}"

	return json, err

	//	return result

}
