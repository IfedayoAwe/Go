package main

// // Encapsulating and appending to a nil slice of string
// type StringStack struct {
// 	data []string
// }

// func (s *StringStack) Push(x string) {
// 	s.data = append(s.data, x)
// }

// func (s *StringStack) Pop() string {
// 	if l := len(s.data); l > 0 {
// 		t := s.data[l-1]
// 		s.data = s.data[:l-1]
// 		return t
// 	}

// 	panic("pop from empty stack")
// }

// // Calling a methon with a nil receiver
// type IntList struct {
// 	value int
// 	Tail  *IntList
// }

// func (list *IntList) Sum() int {
// 	if list == nil {
// 		return 0
// 	}

// 	return list.value + list.Tail.Sum()
// }
