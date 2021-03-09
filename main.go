package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	numOfCounters := 4
	numOfCustomers := 8
	if numOfCounters > numOfCustomers {
		fmt.Println("Number of customers need to be greater than number of counters")
		return
	}

	start := time.Now()
	allocator := newAllocator(numOfCounters).start()
	customers := []int{}
	for i := 1; i <= numOfCustomers; i++ {
		customers = append(customers, i)
	}
	for _, id := range customers {
		allocator.submit(Customer{
			ID:   id,
			Name: fmt.Sprintf("CustomerID: %d", id),
		})
	}
	end := time.Now()
	log.Println("Time taken to serve customers: ", end.Sub(start).Seconds())
}
