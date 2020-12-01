package main

import (
	"fmt"
	"sync"
)

type request struct {
	id int
}

type response struct {
	message string
}

type multResponse struct {
	response []response
}

func ProcessSingle(r request) response {
	return response{
		message: fmt.Sprintf("Processed request %v", r.id),
	}
}

func ProcessMultiple(r []request) multResponse {
	input := make(chan response)
	output := make(chan multResponse)
	defer close(output)
	var wg sync.WaitGroup

	go handleResults(&wg, input, output)

	for _, current := range r {
		wg.Add(1)
		go processConcurrent(current, input)
	}
	wg.Wait()
	close(input)
	result := <-output
	return result
}

func handleResults(wg *sync.WaitGroup, input chan response, output chan multResponse) {
	var result multResponse
	for item := range input {
		result.response = append(result.response, item)
		wg.Done()
	}
	output <- result
}

func processConcurrent(r request, output chan response) {
	resp := ProcessSingle(r)
	output <- resp
}

func main() {

	rmult := []request{
		request{
			id: 1,
		},
		request{
			id: 2,
		},
	}

	for _, current := range rmult {
		resp := ProcessSingle(current)
		fmt.Println(resp)
	}

	result := ProcessMultiple(rmult)
	fmt.Println(result)
}
