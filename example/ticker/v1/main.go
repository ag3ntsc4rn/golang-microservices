package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tick struct {
	ticker    *time.Ticker
	executing bool
}

func somethingYouWantToDo(tick *Tick, t time.Time, processNumber int) {
	fmt.Printf("process#: %v, current time: %v\n", processNumber, t)
	if tick.executing {
		fmt.Printf("process# %v will be skipped. previous run not done yet\n", processNumber)
		return
	}

	tick.executing = true
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 10
	sleepFor := rand.Intn(max-min+1) + min
	fmt.Printf("process# %v triggered and will run for %v seconds\n", processNumber, sleepFor)
	time.Sleep(time.Duration(sleepFor) * time.Second)
	tick.executing = false
}

func main() {
	tick := &Tick{
		ticker: time.NewTicker(5 * time.Second),
	}
	defer tick.ticker.Stop()
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 2000
	for {
		select {
		case t := <-tick.ticker.C:
			go somethingYouWantToDo(tick, t, rand.Intn(max-min)+min)
		}
	}
}
