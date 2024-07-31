
package stress

// import (
// 	"fmt"
// 	"time"
// )

func IterativeFibonacci(n int) int {

	if n <= 1 {
		return n
	}
	fib_min_2 := 0
	fib_min_1 := 1
	fib := 0
	for i := 2; i <= n; i++ {
		fib = fib_min_1 + fib_min_2
		fib_min_2 = fib_min_1
		fib_min_1 = fib
	}
	return fib

}

func RecursiveFibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return RecursiveFibonacci(n-1) + RecursiveFibonacci(n-2)
}

// // example call: go run Go_fibonacci.go
// func main() {
// 	fmt.Print("Enter the number of elements: ")

// 	var fibNum int
// 	fmt.Scanf("%d", &fibNum)

// 	fmt.Print("Fibonacci Series:\n")
// 	for i := 0; i < fibNum; i++ {
// 		nr := time.Now()
// 		fr := RecursiveFibonacci(i)
// 		tr := time.Since(nr)

// 		ni := time.Now()
// 		fi := IterativeFibonacci(i)
// 		ti := time.Since(ni)

// 		fmt.Printf("fib(%d) == %d (%s): %d (%s) \n", i, fr, tr, fi, ti)
// 	}
// }
