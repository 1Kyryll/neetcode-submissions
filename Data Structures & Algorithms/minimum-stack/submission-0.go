type MinStack struct {
	elements []int
}

func Constructor() *MinStack {
	return &MinStack{
		elements: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.elements = append(this.elements, val)
}

func (this *MinStack) Pop() int {
	if len(this.elements) == 0 {
		return 0
	}

	n := len(this.elements) - 1
	data := this.elements[n]

	this.elements[n] = 0
	this.elements = this.elements[:n]

	return data
}

func (this *MinStack) Top() int {
	if len(this.elements) == 0 {
		return 0
	}

	return this.elements[len(this.elements)-1]
}

func (this *MinStack) GetMin() int {
	min := this.Top()
	aux := []int{}

	for len(this.elements) > 0 {
		v := this.Pop()
		if v <= min {
			min = v
		}

		aux = append(aux, v)
	}

	for i := len(aux) - 1; i >= 0; i-- {
		this.Push(aux[i])
	}

	return min
}