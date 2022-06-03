package pushswap

type Stack struct {
	Nums   []int
	length int
}

func NewStack(nums []int, length int) *Stack {
	return &Stack{Nums: nums, length: length}
}

func (s *Stack) Push(n int) {
	s.Nums = append(s.Nums[:s.length], n)
	s.length++
}

func (s *Stack) Pop() int {
	if s.length == 0 {
		return -1
	}
	s.length--
	return s.Nums[s.length]
}

func (s *Stack) PushFront(n int) {
	tmp := []int{n}
	s.Nums = append(tmp, s.Nums[:s.length]...)
	s.length++
}

func (s *Stack) PopFront() int {
	if s.length == 0 {
		return -1
	}
	res := s.Nums[0]
	s.Nums = s.Nums[1:]
	s.length--
	return res
}
