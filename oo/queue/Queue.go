package queue

type Queue []int

func (receiver *Queue) Push(value int) {
	*receiver = append(*receiver, value)
}

func (receiver *Queue) Pop() int {
	head := (*receiver)[0]
	*receiver = (*receiver)[1:]
	return head
}

func (receiver *Queue) IsEmpty() bool {
	return len(*receiver) == 0
}
