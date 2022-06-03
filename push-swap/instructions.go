package pushswap

func pa(s []*Stack) { //push the top first element of stack b to stack a
	if s[1].length > 0 {
		s[0].Push(s[1].Pop())
	}
}

func pb(s []*Stack) { //push the top first element of stack a to stack b
	if s[0].length > 0 {
		s[1].Push(s[0].Pop())
	}
}

func sa(s []*Stack) { //swap first 2 elements of stack a
	if s[0].length > 1 {
		v1, v2 := s[0].Pop(), s[0].Pop()
		s[0].Push(v1)
		s[0].Push(v2)
	}
}
func sb(s []*Stack) { //swap first 2 elements of stack b
	if s[1].length > 1 {
		v1, v2 := s[1].Pop(), s[1].Pop()
		s[1].Push(v1)
		s[1].Push(v2)
	}
}
func ss(s []*Stack) { //execute sa and sb
	sa(s)
	sb(s)
}
func ra(s []*Stack) { //rotate stack a (shift up all elements of stack a by 1, the first element becomes the last one)
	if s[0].length == 0 {
		return
	}
	s[0].PushFront(s[0].Pop())
}
func rb(s []*Stack) { //rotate stack b
	if s[1].length == 0 {
		return
	}
	s[1].PushFront(s[1].Pop())
}
func rr(s []*Stack) { //execute ra and rb
	ra(s)
	rb(s)
}
func rra(s []*Stack) { //reverse rotate a (shift down all elements of stack a by 1, the last element becomes the first one)
	if s[0].length == 0 {
		return
	}
	s[0].Push(s[0].PopFront())
}
func rrb(s []*Stack) { //reverse rotate b
	if s[1].length == 0 {
		return
	}
	s[1].Push(s[1].PopFront())
}
func rrr(s []*Stack) { //execute rra and rrb
	rra(s)
	rrb(s)
}
