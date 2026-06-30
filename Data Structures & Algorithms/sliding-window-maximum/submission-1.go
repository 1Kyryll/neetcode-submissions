func maxSlidingWindow(nums []int, k int) []int {
	if k > len(nums) || k == 0 {
		return []int{}
	}

	heap := &MaxHeap{array: make([]Num, 0, len(nums))}
	res := make([]int, 0, len(nums)-k+1)

	for i := 0; i < len(nums); i++ {
		heap.Insert(Num{index: i, value: nums[i]})
		
		if i >= k-1 {
			for {
				root, _ := heap.Peek()
				if root.index < i-k+1 {
					heap.Extract()
				} else {
					break
				}
			}

			maxNode, _ := heap.Peek()
			res = append(res, maxNode.value)
		}
	}

	return res
}

type Num struct {
	index int
	value int
}

type MaxHeap struct {
	array []Num
}

func (h *MaxHeap) Extract() (Num, bool) {
	if len(h.array) == 0 {
		return Num{}, false
	}

	max := h.array[0]
	lastIdx := len(h.array) - 1
	h.array[0] = h.array[lastIdx]
	h.array = h.array[:lastIdx]

	if len(h.array) > 0 {
		h.heapifyDown(0)
	}

	return max, true
}

func (h *MaxHeap) Peek() (Num, bool) {
	if len(h.array) == 0 {
		return Num{}, false
	}

	return h.array[0], true
}

func (h *MaxHeap) Insert(key Num) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parentIdx := (index - 1) / 2
		if h.array[parentIdx].value >= h.array[index].value {
			break
		}
		h.array[index], h.array[parentIdx] = h.array[parentIdx], h.array[index]
		index = parentIdx
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	lastIdx := len(h.array) - 1
	l, r := 2*index+1, 2*index+2
	childToCompare := 0

	for l <= lastIdx {
		if l == lastIdx {
			childToCompare = l
		} else if h.array[l].value > h.array[r].value {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index].value > h.array[childToCompare].value {
			break
		}

		h.array[index], h.array[childToCompare] = h.array[childToCompare], h.array[index]
		index = childToCompare
		l, r = 2*index+1, 2*index+2
	}
}

