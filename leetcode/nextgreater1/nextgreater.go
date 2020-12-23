package nextgreater

type Stack struct {
	items []int
}

func Constructor() Stack {
	return Stack{
		items: []int{},
	}
}
func (this *Stack) Push(x int) {
	this.items = append(this.items, x)
}
func (this *Stack) Pop() int {
	if len(this.items) == 0 {
		panic("Empty Stack")
	}
	index := len(this.items) - 1
	item := this.items[index]
	this.items = this.items[:index]
	return item
}
func (this *Stack) IsEmpty() bool {
	return len(this.items) == 0
}
func (this *Stack) Peek() int {
	return this.items[len(this.items)-1]
}
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextMap := map[int]int{}
	stack := Constructor()
	for _, num := range nums2 {
		for !stack.IsEmpty() && stack.Peek() < num {
			nextMap[stack.Pop()] = num
		}
		stack.Push(num)
	}
	for i, num := range nums1 {
		if v, ok := nextMap[num]; ok {
			nums1[i] = v
		} else {
			nums1[i] = -1
		}
	}
	return nums1
}
