package tgbot

import (
	"fmt"
	"time"
)

type Poller interface {
	Listen(Pollable, chan Update, chan struct{})
}

type Pollable interface {
	GetUpdates(UpdateId) ([]Update, error)
}

type LongPoller struct {
	lastUpdateId UpdateId
}

func NewPoller() *LongPoller {
	return &LongPoller{}
}

func (p *LongPoller) Listen(source Pollable, dest chan Update, stop chan struct{}) {
	fmt.Println("Listen to updates")
	const TimeInterval float64 = 5.0
	start := time.Now()
	for {
		select {
		case <- stop:
			return
		default:
		}

		if time.Since(start).Seconds() >= TimeInterval {
			fmt.Println("Trying to get updates...")
			updates, _ := source.GetUpdates(p.lastUpdateId + 1)
			for _, update := range updates {
				p.lastUpdateId = update.UpdateId
				dest <- update
			}
			start = time.Now()
		}
	}
}
