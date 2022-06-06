package utils

type Stack struct {
	Nums   []int
	Length int
}

func NewStack(nums []int, Length int) *Stack {
	return &Stack{Nums: nums, Length: Length}
}

func (s *Stack) Push(n int) {
	s.Nums = append(s.Nums[:s.Length], n)
	s.Length++
}

func (s *Stack) Pop() int {
	if s.Length == 0 {
		return -1
	}
	s.Length--
	return s.Nums[s.Length]
}

func (s *Stack) PushFront(n int) {
	tmp := []int{n}
	s.Nums = append(tmp, s.Nums[:s.Length]...)
	s.Length++
}

func (s *Stack) PopFront() int {
	if s.Length == 0 {
		return -1
	}
	res := s.Nums[0]
	s.Nums = s.Nums[1:]
	s.Length--
	return res
}
