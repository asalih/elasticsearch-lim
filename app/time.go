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
	t.ticker = time.NewTicker(time.Minute * 2)
	if current != nil {
		current.Stop()
	}

	current = t.ticker

	func() {
		elastic := &ElasticHandler{}
		fmt.Println("ticker")
		for t := range t.ticker.C {
			elastic.Time = t
			elastic.CollectNewData()
			fmt.Println(t)
		}
	}()

}

func GetTimer() *time.Ticker {
	return current
}
