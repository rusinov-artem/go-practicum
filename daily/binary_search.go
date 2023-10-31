package daily

func search(slice []int, val int) int {
	if len(slice) < 1 {
		return -1
	}

	begin := 0
	end := len(slice)
	var mid int

	for begin < end {
		mid = begin + (end-begin)/2

		if slice[mid] < val {
			begin = mid + 1
		} else if slice[mid] > val {
			end = mid
		} else {
			break
		}

	}

	if slice[mid] == val {
		return mid
	}

	return -1
}

func lowerBound(slice []int, val int) int {
	begin := 0
	end := len(slice)
	var mid int

	for begin < end {
		mid = begin + (end-begin)/2
		if slice[mid] < val {
			begin = mid+1
		} else if slice[mid] >= val {
			end = mid
		}
	}

	return begin
}

func upperBound(slice []int, val int) int {
	begin := 0
	end := len(slice)
	var mid int

	for begin < end {
		mid = begin + (end-begin)/2
		if slice[mid] <= val {
			begin = mid+1
		} else if slice[mid] > val {
			end = mid
		}
	}

	return end
}

