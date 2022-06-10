package pushswap

import (
	"pushswap/utils"
)

func pa(s []*utils.Stack) { //push the top first element of stack b to stack a
	if len(s[1].Nums) > 0 {
		s[0].PushBack(s[1].PopBack())
	}
}

func pb(s []*utils.Stack) { //push the top first element of stack a to stack b
	if len(s[0].Nums) > 0 {
		s[1].PushBack(s[0].PopBack())
	}
}

func sa(s []*utils.Stack) { //swap first 2 elements of stack a
	if len(s[0].Nums) > 1 {
		v1 := s[0].PopBack()
		v2 := s[0].PopBack()
		s[0].PushBack(v1)
		s[0].PushBack(v2)
	}
}
func sb(s []*utils.Stack) { //swap first 2 elements of stack b
	if len(s[1].Nums) > 1 {
		v1 := s[1].PopBack()
		v2 := s[1].PopBack()
		s[1].PushBack(v1)
		s[1].PushBack(v2)
	}
}
func ss(s []*utils.Stack) { //execute sa and sb
	sa(s)
	sb(s)
}
func ra(s []*utils.Stack) { //rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
	if len(s[0].Nums) == 0 {
		return
	}
	s[0].PushFront(s[0].PopBack())
}
func rb(s []*utils.Stack) { //rotate stack b
	if len(s[1].Nums) == 0 {
		return
	}
	s[1].PushFront(s[1].PopBack())
}
func rr(s []*utils.Stack) { //execute ra and rb
	ra(s)
	rb(s)
}
func rra(s []*utils.Stack) { //reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
	if len(s[0].Nums) == 0 {
		return
	}
	s[0].PushBack(s[0].PopFront())
}
func rrb(s []*utils.Stack) { //reverse rotate b
	if len(s[1].Nums) == 0 {
		return
	}
	s[1].PushBack(s[1].PopFront())
}
func rrr(s []*utils.Stack) { //execute rra and rrb
	rra(s)
	rrb(s)
}
