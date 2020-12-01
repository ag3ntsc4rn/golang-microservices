package main

import (
	"fmt"
	"sync"
	"time"
)

type SomeApiRequest struct {
	Name string
}

type SomeApiResponse struct {
	Message string
}

type SomeApiResponseCollection struct {
	Results []SomeApiResponse
}

func ProcessRequest(request SomeApiRequest) SomeApiResponse {
	fmt.Printf("Processing Request #: %v\n", request.Name)
	time.Sleep(10 * time.Second)
	return SomeApiResponse{
		Message: fmt.Sprintf("Response for: %s", request.Name),
	}
}

func main() {
	var requests []SomeApiRequest
	var response []SomeApiResponse
	for i := 0; i < 10; i++ {
		r := SomeApiRequest{
			Name: fmt.Sprintf("Request #%d", i),
		}
		requests = append(requests, r)
	}

	var wg sync.WaitGroup
	input := make(chan SomeApiResponse)
	output := make(chan []SomeApiResponse)
	buffer := make(chan bool, 2)
	defer close(output)

	go handleConcurrentResponse(&wg, input, output)
	for _, request := range requests {
		// response = append(response, ProcessRequest(request))
		wg.Add(1)
		buffer <- true
		go concurrentApiRequests(request, input, buffer)
	}
	wg.Wait()
	close(input)
	response = <-output
	fmt.Println(len(response))
}

func handleConcurrentResponse(wg *sync.WaitGroup, input chan SomeApiResponse, output chan []SomeApiResponse) {
	var response []SomeApiResponse
	for incomingEvent := range input {
		response = append(response, incomingEvent)
		wg.Done()
	}
	output <- response
}

func concurrentApiRequests(request SomeApiRequest, output chan SomeApiResponse, buffer chan bool) {
	r := ProcessRequest(request)
	<-buffer
	output <- r
}
