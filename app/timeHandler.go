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
	t.ticker = time.NewTicker(time.Millisecond * 1000)
	if current != nil {
		current.Stop()
	}

	current = t.ticker

	func() {
		for t := range t.ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

}

func GetTimer() *time.Ticker {
	return current
}
