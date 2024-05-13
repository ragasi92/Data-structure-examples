package main

import "log"

//left child index = (parent index x 2) + 1

//right child index = (parent index x 2) + 2

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	if len(h.array) == 0 {
		log.Println("The heap is empty")
		return -1
	}
	lastelementIndex := len(h.array) - 1
	h.array[0] = h.array[lastelementIndex]
	h.array = h.array[:lastelementIndex]
	h.maxHeapifyDown(0)
	return extracted
}

func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[getParentIndex(index)] < h.array[index] {
		h.swap(getParentIndex(index), index)
		index = getParentIndex(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastElementIndex := len(h.array) - 1
	left, right := getLeftChild(index), getRightChild(index)
	childToCompare := 0

	for left <= lastElementIndex {
		if left == lastElementIndex {
			childToCompare = left
		} else if h.array[left] > h.array[right] {
			childToCompare = left
		} else {
			childToCompare = right
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			left, right = getLeftChild(index), getRightChild(index)
		} else {
			return
		}
	}

}

func getParentIndex(i int) int {
	return (i - 1) / 2
}

func getLeftChild(parentIndex int) int {
	return (parentIndex * 2) + 1
}
func getRightChild(parentIndex int) int {
	return (parentIndex * 2) + 2
}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	log.Println("=--=-=-=-==-=--==main-==--==--=")

	maxheap := &MaxHeap{}
	buildHeap := []int{1, 3, 6, 5, 4, 9, 30, 60, 17, 14, 2}

	for _, value := range buildHeap {
		maxheap.Insert(value)
		log.Println(maxheap)
	}

	for i := 0; i < 6; i++ {
		maxheap.Extract()
		log.Println(maxheap)
	}
}
