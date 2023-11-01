package main

func brokenSearch(arr []int, k int) int {
	begin := 0
	end := len(arr)
	var mid int

	for begin < end {
		mid = begin + (end-begin)/2
		if arr[mid] == k {
			return mid
		}
		if arr[begin] < arr[mid] {
			if k >= arr[begin] && k <= arr[mid] {
				end = mid
			} else {
				begin = mid + 1
			}
		} else {
			if k <= arr[end-1] && k >= arr[mid] {
				begin = mid + 1
			} else {
				end = mid
			}
		}
	}
	return -1
}
