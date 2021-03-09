# Queue Processing System

Queue processing system is a system that consists of N queues and X customers. It processes the X customers in minimum time possible

### HOW TO RUN THE SCRIPT
* Type "go run ." without the double quotes while navigating to the counter directory

### HOW TO RUN THE TESTCASE
* Type "go test -bench=." without the double quotes while navigating to the current directory

NOTE: You can change the variable numOfCustomers, numOfCounters in main.go and numCustomers, numCounters in main_test.go just to play around.

### APPROACH
 - The Allocator starts "n" number of counters/workers.
 - The multiple workers share a common queue.
 - The allocator submites a job or a customer on the WorkChannel.
 - When the job is availabe on the workChannel, the counter or the worker pushes it into the queue.
 - When the job is available, the counter/worker processes it and moves on to processing the next job whenever it is available on the work channel.
 - This process continues until all the customers have been served.
