package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	lock sync.Mutex
	atomicInt = AtomicInt{}
)

type AtomicInt struct{
	value int
	lock sync.Mutex
}

func (i *AtomicInt) Increase(){
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value++	
}

func (i *AtomicInt) Value() int{
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.value
}

func main() {
	var wg sync.WaitGroup
	for i:=0; i<10000; i++{
		wg.Add(1)
		go incrementCounter(&wg)
	}
	wg.Wait()
	fmt.Printf("Counter value is %d\n", counter)
	fmt.Printf("Atomic counter value is %d\n", atomicInt.Value())
	
}

func incrementCounter(wg *sync.WaitGroup) {
	lock.Lock()
	defer lock.Unlock()
	counter++
	atomicInt.Increase()
	wg.Done()
}
