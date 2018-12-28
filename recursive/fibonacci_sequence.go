package recursive

// 0,1,1,2,3,5,8,13,21,34,55
// 2 = f(1) + f(0)
// 3 = f(2) + f(1)
// 4 = f(3) + f(2)
// f(n) = f(n-1)+f(n-2)
// f(n) = f(n-2) + f(n-3)

func Fibonacci(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

var m = map[uint]uint{
	0: 0,
	1: 1,
}

func FibonacciWithMap(n uint) uint {
	if n < 2 {
		return m[n]
	}
	if _, ok := m[n]; !ok {
		m[n] = FibonacciWithMap(n-1) + FibonacciWithMap(n-2)
	}
	return m[n]
}
