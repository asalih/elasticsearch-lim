package app

import (
	"fmt"
	"time"
)

var current *time.Ticker

type TimeHandler struct {
	ticker *time.Ticker
}

func (t *TimeHandler) InitTime() {
	t.ticker = time.NewTicker(time.Second * 2)
	if current != nil {
		current.Stop()
	}

	current = t.ticker

	func() {
		elastic := &ElasticHandler{}
		for t := range t.ticker.C {
			elastic.CollectNewData()
			fmt.Println("Tick at", t)
		}
	}()

}

func GetTimer() *time.Ticker {
	return current
}
