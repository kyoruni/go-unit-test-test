package main

// 偶数か奇数か返す
func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "even"
	}
	return "odd"
}
