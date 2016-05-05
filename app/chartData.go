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
	result := req.DoRawPostRequest(req.CT(os.Getenv("ELASTICSEARCH_INDEX_"+env)+"_search"), "text/json", qs)

	return result

}

func (c *ChartData) GetLoadData(unixT int64) string {
	qt := &q.QueryTemplates{}
	qs := qt.GetIndicesTemplate(unixT)

	req := &RequestHandler{}
	result := req.DoRawPostRequest(req.CT(os.Getenv("ELASTICSEARCH_INDEX_ST")+"_search"), "text/json", qs)
	resultN := req.DoRawPostRequest(req.CT(os.Getenv("ELASTICSEARCH_INDEX_ND")+"_search"), "text/json", qs)

	return "{\"s\": " + result + ", \"n\": " + resultN + "}"
	//	return result

}
