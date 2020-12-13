package main

import (
	"fmt"
	"math/rand"
	"time"
)

var inProgress bool

func main() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		rand.Seed(time.Now().UnixNano())
		min := 1000
		max := 2000
		select {
		case t := <-ticker.C:
			p := rand.Intn(max-min) + min
			fmt.Printf("Tick at %v. Will try to run process# %v\n", t, p)
			if !inProgress {
				inProgress = true
				go temp(t, rand.Intn(max-min)+min)
			} else {
				fmt.Printf("Skip process# %v\n", p)
			}
		}
	}
}

func temp(t time.Time, process int) {
	defer func() {
		inProgress = false
	}()
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 10
	sleepFor := rand.Intn(max-min+1) + min
	fmt.Printf("process# %v triggered and will run for %v seconds\n", process, sleepFor)
	time.Sleep(time.Second * time.Duration(sleepFor))
}
