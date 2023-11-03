package daily

func partition(d []int, begin, end int) int {
	if end-begin < 2 {
		return begin
	}

	left := begin
	right := end - 2
	mid := end - 1

	for {
		for d[mid] >= d[left] && left < mid {
			left++
		}

		for d[right] > d[mid] && right > begin {
			right--
		}

		if left < right {
			swap(d, left, right)
		} else {
			break
		}
	}

	swap(d, mid, left)

	return left
}

func swap(d []int, i, j int) {
	d[i], d[j] = d[j], d[i]
}

func qSort(d []int) {
	qsort(d, 0, len(d))
}

func qsort(d []int, begin, end int) {
	if end-begin <2 {
		return 
	}
	p := partition(d, begin, end)
	qsort(d, begin, p)
	qsort(d, p+1, end)
}
