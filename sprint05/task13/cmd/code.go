package main

func siftUp(heap []int, idx int) int {
	return up(heap, idx)
}

func up(heap []int, idx int) int {
	for p := parent(idx); idx > 1; p=parent(idx) {
		if heap[p] < heap[idx] {
			swap(heap, idx, p)
			idx = p
		} else {
           return idx
        }
	}
	return idx
}

func parent(idx int) int {
	return idx / 2
}

func leftChild(idx int) int {
	return idx * 2
}

func rightChild(idx int) int {
	return idx*2 + 1
}

func swap(heap []int, l, r int) {
	heap[l], heap[r] = heap[r], heap[l]
}
