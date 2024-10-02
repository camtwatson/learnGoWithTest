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

func sumAll(sumNums ...[]int) []int {
	var sums []int

	for _, allnums := range sumNums {
		sums = append(sums, sumSlice(allnums))
	}
	return sums
}

func sumAllTails(sumNums ...[]int) []int {
	var sums []int

	for _, tailnums := range sumNums {
		if len(tailnums) == 0 {
			sums = append(sums, 0)
		} else {
			tail := tailnums[1:]
			sums = append(sums, sumSlice(tail))
		}
	}
	return sums
}
