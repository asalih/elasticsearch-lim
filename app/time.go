package app

import (
	"os"
	"strconv"
	"time"
)

var current *time.Ticker

type TimeHandler struct {
	ticker *time.Ticker
}

func (t *TimeHandler) InitTime() {
	intvSec, _ := strconv.Atoi(os.Getenv("INTERVAL_SECOND"))
	t.ticker = time.NewTicker(time.Duration(intvSec) * time.Second)
	if current != nil {
		current.Stop()
	}

	current = t.ticker

	elastic := &ElasticHandler{}

	for t := range t.ticker.C {
		elastic.Time = t
		elastic.CollectNewData()
	}

}

func GetTimer() *time.Ticker {
	return current
}
