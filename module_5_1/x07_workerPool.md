Workers pool
One problem we may face with some of the previous approaches to concurrency 
is their unbounded context. We cannot let an app create an unlimited amount of Goroutines.
Goroutines are light, but the work they perform could be very heavy. A workers pool helps us to solve this problem.
