package main

// Sum 计算切片 numbers 的元素的总和
func Sum(numbers []int) int {
	sum := 0
	// 使用 range，那么遍历的就是 index，value
	for _, numbers := range numbers {
		sum += numbers
	}
	return sum
}

// SumAll 计算numbers（多个切片）的总和，然后按照顺序存储到 sums 数组之中。
func SumAll(numbers ...[]int) []int {
	lengthOfNumbers := len(numbers)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbers {
		sums[i] = Sum(numbers)
	}

	return sums
}
