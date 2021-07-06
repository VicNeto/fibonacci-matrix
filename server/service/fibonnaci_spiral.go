package service

func FibonnaciSpiral(rows, cols int) ([][]int, error) {
	sequence := fibo(rows * cols)
	spiral := [][]int{}

	for i := 0; i < rows; i++ {
		spiral = append(spiral, sequence[cols*i:cols*(i+1)])
	}
	return spiral, nil
}

func fibo(n int) []int {
	fib := []int{1}

	if n == 1 {
		return fib
	}
	fib = append(fib, 1)

	for i := 2; i < n; i++ {
		fib = append(fib, fib[len(fib)-1]+fib[len(fib)-2])
	}
	return fib
}
