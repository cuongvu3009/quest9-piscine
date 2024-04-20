package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) > 1 {
		if f(a[0], a[1]) >= 0 {
			for i := 0; i < len(a)-1; i++ {
				if f(a[i], a[i+1]) < 0 {
					return false
				}
			}
		}
		if f(a[0], a[1]) <= 0 {
			for i := 0; i < len(a)-1; i++ {
				if f(a[i], a[i+1]) > 0 {
					return false
				}
			}
		}
	}
	return true
}
func f(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
