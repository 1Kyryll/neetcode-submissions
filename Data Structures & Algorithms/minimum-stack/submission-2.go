type MinStack struct {
	elements []int
	minS     []int
}

func Constructor() *MinStack {
	return &MinStack{
	elements: make([]int, 0),
	minS:     make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.elements = append(this.elements, val)
	if len(this.minS) == 0 {
		this.minS = append(this.minS, val)
		return
	}

	topMin := this.minS[len(this.minS)-1]
	if val <= topMin {
		this.minS = append(this.minS, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.elements) == 0 {
		return
	}

	n := len(this.elements) - 1
	data := this.elements[n]

	this.elements = this.elements[:n]
	if data == this.minS[len(this.minS)-1] {
		this.minS = this.minS[:len(this.minS)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.elements) == 0 {
		return 0
	}

	return this.elements[len(this.elements)-1]
}

func (this *MinStack) GetMin() int {
	return this.minS[len(this.minS)-1]
}
