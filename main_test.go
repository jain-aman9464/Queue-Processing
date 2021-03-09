package main

import (
	"fmt"
	"testing"
)

const (
	numCustomers = 10
	numCounters  = 4
)

func ConcurrentServicing(b *testing.B) {
	if numCounters > numCustomers {
		fmt.Println("Number of customers need to be greater than number of counters")
		return
	}
	custs := []int{}
	for i := 1; i <= numCustomers; i++ {
		custs = append(custs, i)
	}
	allocator := newAllocator(numCounters).start()
	for n := 0; n < b.N; n++ {
		for i := range custs {
			allocator.submit(Customer{
				ID:   i,
				Name: fmt.Sprintf("CustomerID: %d", i),
			})
		}
	}
}
