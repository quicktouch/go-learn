package series

import "fmt"

func GetFibonacciSeries(n int) ([]int, error) {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func getFibonacciSeries(n int) ([]int, error) {
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func init() {
	fmt.Println("init-series-1")
}

func init() {
	fmt.Println("init-series-2")
}
