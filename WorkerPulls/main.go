package main

import "fmt"

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))
	nWorkers := 4
	for i := 1; i <= nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, item := range tasks {
		jobs <- item
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}

func Worker(idWorker int, jobs <-chan int, results chan<- int) {
	for value := range jobs {

		fmt.Printf("Trabajador actual id: %d trabajando en el job: %d\n", idWorker, value)
		fib := Fibonacci(value)
		fmt.Printf("Finalización trabajador id: %d trabajando en el job: %d resultado: %d\n", idWorker, value, fib)

		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
