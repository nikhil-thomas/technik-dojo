package calc

// Sum returns the sum of elements of ints
func Sum(vals []int) int {
	var sum int
	for _, val := range vals {
		sum += val
	}
	return sum
}

// SumAll find sums of slices
func SumAll(valSlices ...[]int) []int {
	result := make([]int, len(valSlices))
	for _, vals := range valSlices {
		sum := Sum(vals)
		result = append(result, sum)
	}
	return result
}

// SumTails find sums of slice tails
func SumTails(valSlices ...[]int) []int {
	result := make([]int, len(valSlices))
	for _, vals := range valSlices {
		var sum int
		if len(vals) != 0 {
			tail := vals[1:]
			sum = Sum(tail)
		}
		result = append(result, sum)
	}
	return result

}
