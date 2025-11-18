package main

func Sum(numbers []int) int {
	sum := 0
	// 使用 range，那么遍历的就是 index，value
	for _, numbers := range numbers {
		sum += numbers
	}
	return sum
}
