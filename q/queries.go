package q

import "fmt"

type QueryTemplates struct {
}

func (q *QueryTemplates) GetCurrentQueryTemplate(unix int64, idx string) string {
	return fmt.Sprintf(`{
  "size": 30,
  "query": {
    "filtered": {
      "query": {
        "match": {
          "idx": "%s"
        }
      },
      "filter": {
        "range": {
          "timestamp": {
            "gte": %d
          }
        }
      }
    }
  }, 
  "sort": [
    {
      "timestamp": {
        "order": "desc"
      }
    }
  ]
}
`, idx, unix)
}

func (q *QueryTemplates) GetIndicesTemplate(unix int64) string {
	return fmt.Sprintf(`{
  "size": 0,
  "query": {
    "filtered": {
      "query": {
        "match_all": {}
      },
      "filter": {
        "range": {
          "timestamp": {
            "gte": %d
          }
        }
      }
    }
  },
  "aggs": {
    "idx_agg": {
      "terms": {
        "field": "idx"
      }
    }
  }
}
`, unix)
}

func (q *QueryTemplates) GetTargetMappingTemplate(ttl string) string {
	return fmt.Sprintf(`{
   "_ttl" : { "enabled" : true, "default" : "%s" }
}
`, ttl)
}
