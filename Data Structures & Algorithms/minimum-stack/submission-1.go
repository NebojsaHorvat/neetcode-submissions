type MinStack struct {
	lenght int
	data []int
	minimum int
}

func Constructor() MinStack {
	return MinStack{
		data:make([]int,0),
	}
}

func (this *MinStack) Push(val int) {
	if this.lenght == 0 {
		this.minimum = val
	}else {
		if val < this.minimum{
			this.minimum = val
		}
	}
	this.data = append(this.data, val)
	this.lenght++
}

func (this *MinStack) Pop() {
	if this.lenght == 0{
		return 
	}
	ret := this.data[this.lenght-1]
	this.data = this.data[:this.lenght-1]
	this.lenght--
	if this.lenght > 0 && ret == this.minimum{
		this.minimum = this.data[0]
		for _, val := range this.data{
			if val < this.minimum {
				this.minimum = val
			}
		}
		
	}
}

func (this *MinStack) Top() int {
	if this.lenght == 0{
		return 0
	}
	return this.data[this.lenght-1]
}

func (this *MinStack) GetMin() int {
	return this.minimum
}
