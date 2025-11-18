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
func SumAll(numbersToSum ...[]int) []int {
	//lengthOfNumbers := len(numbers)
	//
	////创建一个容量为 lengthOfNumbers 的切片。
	////length: 长度则是已经是用了的容量，或者说已存储元素的数量
	////capacity: 容量就是可以容纳的元素的数量
	//sums := make([]int, lengthOfNumbers)
	//
	//for i, numbers := range numbers {
	//	sums[i] = Sum(numbers)
	//}
	//
	//return sums

	// refactor
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers)) // 接受一个切片，一个新值。随后返回包含所有元素的新切片。
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
