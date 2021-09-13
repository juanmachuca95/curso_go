package main

/* import "fmt" */

func Sum(a, b int) int {
	return a + b
}

func GetMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Fibonacci(x int) int {
	if(x <= 1){
		return x;
	}
	return Fibonacci(x - 1) + Fibonacci(x - 2)
}

/* func main() {
	fmt.Printf("%d\n", Sum(1, 2))
} */
