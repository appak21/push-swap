package utils

type Stack struct {
	Nums []int
}

func NewStack(nums []int) *Stack {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
	return &Stack{nums}
}

func (s *Stack) PushBack(n int) {
	s.Nums = append(s.Nums, n)
}

func (s *Stack) PopBack() int {
	if len(s.Nums) == 0 {
		return -1
	}
	res := s.Nums[len(s.Nums)-1]
	s.Nums = s.Nums[:len(s.Nums)-1]
	return res
}

func (s *Stack) PushFront(n int) {
	s.Nums = append([]int{n}, s.Nums...)
}

func (s *Stack) PopFront() int {
	if len(s.Nums) == 0 {
		return -1
	}
	res := s.Nums[0]
	s.Nums = s.Nums[1:]
	return res
}
