package main

func merge_sort(arr []int, lf int, rg int) {
	if rg-lf < 2 {
		return
	}
	mid := lf + (rg-lf)/2
	merge_sort(arr, lf, mid)
	merge_sort(arr, mid, rg)
	m := merge(arr, lf, mid, rg)
	for i := lf; i < rg; i++ {
		arr[i] = m[i-lf]
	}

}

func merge(arr []int, left int, mid int, right int) []int {
	res := make([]int, 0, right-left)

	lb := left
	le := mid

	rb := mid
	re := right

	for lb < le || rb < re {
		if lb >= le {
			res = append(res, arr[rb])
			rb++
		} else if rb >= re {
			res = append(res, arr[lb])
			lb++
		} else if arr[lb] > arr[rb] {
			res = append(res, arr[rb])
			rb++
		} else {
			res = append(res, arr[lb])
			lb++
		}
	}

	return res
}

// func main() {
// 	a := []int{1, 4, 9, 2, 10, 11}
// 	b := merge(a, 0, 3, 6)
// 	expected := []int{1, 2, 4, 9, 10, 11}
// 	if !reflect.DeepEqual(b, expected) {
// 		panic("WA. Merge")
// 	}

// 	c := []int{1, 4, 2, 10, 1, 2}
// 	merge_sort(c, 0, 6)
// 	expected = []int{1, 1, 2, 2, 4, 10}
// 	if !reflect.DeepEqual(c, expected) {
// 		panic("WA. MergeSort")
// 	}
// }
