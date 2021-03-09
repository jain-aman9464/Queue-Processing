package main

import (
	"fmt"
	"time"
)

const (
	timeTakenToServeOneCustomer = 500
)

// Job represents a single customer
type Customer struct {
	ID   int
	Name string
}

type CustomerChannel chan Customer
type CustomerQueue chan chan Customer

// A Counter serves a particular customer at any point of time
type Counter struct {
	ID           int             // id of the counter
	CustomerChan CustomerChannel // a channel to receive single unit of work
	Queue        CustomerQueue   // shared between all counters.
	Quit         chan struct{}   // a channel to quit working
}

func newCounter(ID int, CustomerChan CustomerChannel, Queue CustomerQueue, Quit chan struct{}) *Counter {
	return &Counter{
		ID:           ID,
		CustomerChan: CustomerChan,
		Queue:        Queue,
		Quit:         Quit,
	}
}

func (cr *Counter) start() {
	go func() {
		for {
			// when available, put the CustomerChan on the Queue
			cr.Queue <- cr.CustomerChan
			select {
			case customer := <-cr.CustomerChan:
				// when a customer is received, process it
				// time.Sleep is basically a simulation of processing of a particular customer
				// I've assumed that it takes 500 milliseconds to serve 1 customer
				fmt.Println("serving customer -> ", customer.ID)
				time.Sleep(timeTakenToServeOneCustomer * time.Millisecond)
			case <-cr.Quit:
				// a signal on this channel means someone triggered
				// a shutdown for this worker
				close(cr.CustomerChan)
				return
			}
		}
	}()
}
