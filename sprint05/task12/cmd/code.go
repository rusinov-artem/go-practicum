package main

func siftDown(heap []int, idx int) int {
	return down(heap, idx)
}

func down(heap []int, idx int) int {
	left := leftChild(idx)
	right := rightChild(idx)

	if left >= len(heap) {
		return idx
	}

	maxChildIdx := left
	if right < len(heap) && heap[right] > heap[left] {
		maxChildIdx = right
	}

	if heap[maxChildIdx] > heap[idx] {
		swap(heap, idx, maxChildIdx)
		return down(heap, maxChildIdx)
	}

	return idx
}

func swap(heap []int, l, r int) {
	heap[l], heap[r] = heap[r], heap[l]
}

func parent(p int) int {
	return p / 2
}

func leftChild(p int) int {
	return p*2 
}

func rightChild(p int) int {
	return p*2 + 1
}
