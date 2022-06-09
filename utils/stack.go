package utils

type Stack struct {
	Nums []int
}

func NewStack(nums []int) *Stack {
	return &Stack{nums}
}

func (s *Stack) Push(n int) {
	s.Nums = append(s.Nums, n)
}

func (s *Stack) Pop() int {
	if len(s.Nums) == 0 {
		return -1
	}
	res := s.Nums[len(s.Nums)-1]
	s.Nums = s.Nums[:len(s.Nums)-1]
	return res
}

func (s *Stack) PushFront(n int) {
	tmp := []int{n}
	s.Nums = append(tmp, s.Nums...)
}

func (s *Stack) PopFront() int {
	if len(s.Nums) == 0 {
		return -1
	}
	res := s.Nums[0]
	s.Nums = s.Nums[1:]
	return res
}
