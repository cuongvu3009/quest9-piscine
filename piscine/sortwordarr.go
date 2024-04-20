package piscine

func SortWordArr(a []string) {
	sort(a, 0, len(a)-1)
}
func sort(table []string, start int, end int) {
	if start < end {
		index := checkAscii(table, start, end)
		sort(table, start, index-1)
		sort(table, index+1, end)
	}
}
func checkAscii(table []string, start int, end int) int {
	lindex := table[end]
	i := start - 1
	for j := start; j < end; j++ {
		if table[j] <= lindex {
			i++
			tmp := table[j]
			table[j] = table[i]
			table[i] = tmp
		}
	}
	tmp := table[end]
	table[end] = table[i+1]
	table[i+1] = tmp
	return i + 1
}
