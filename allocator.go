package main

// New returns a new allocator
func newAllocator(num int) *Allocator {
	return &Allocator{
		Counters: make([]*Counter, num),
		WorkChan: make(CustomerChannel),
		Queue:    make(CustomerQueue),
	}
}

type Allocator struct {
	Counters []*Counter      // this is the list of counters
	WorkChan CustomerChannel // Job is submitted to this channel
	Queue    CustomerQueue   // this is the shared JobPool between the multiple counters
}

// Start creates pool of counters of count = num
func (atr *Allocator) start() *Allocator {
	l := len(atr.Counters)
	for i := 1; i <= l; i++ {
		ct := newCounter(i, make(CustomerChannel), atr.Queue, make(chan struct{}))
		ct.start()
		atr.Counters = append(atr.Counters, ct)
	}
	go atr.process()
	return atr
}

// process listens to a job submitted on WorkChan and
func (atr *Allocator) process() {
	for {
		select {
		case job := <-atr.WorkChan: // listen to any submitted job on the WorkChan
			jobChan := <-atr.Queue
			jobChan <- job
		}
	}
}

func (atr *Allocator) submit(job Customer) {
	atr.WorkChan <- job
}
