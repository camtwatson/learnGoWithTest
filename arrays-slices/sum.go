package main

func sum(nums [5]int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func sumSlice(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
