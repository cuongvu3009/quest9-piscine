package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, num := range a {
		result[i] = f(num)
	}
	return result
}
