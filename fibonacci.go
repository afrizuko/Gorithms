package gorithms

//fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2)
func fib() func() int {
	a, b := 0, 1
	return func() int {
		defer func() { a, b = b, a+b }()
		return a
	}
}
