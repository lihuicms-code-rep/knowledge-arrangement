package main

import "fmt"

type Student struct {
	name string
}

func NewStudent(name string) *Student {
	return &Student{
		name:name,
	}
}

func (s *Student) update(name string) {
	fmt.Printf("before s:%+v, addr:%p, s addr:%p\n", s, s, &s)
	s = NewStudent(name)
	fmt.Printf("after s:%+v, addr:%p, s addr:%p\n", s, s, &s)
}

func main() {
	s1 := NewStudent("lihui")
	fmt.Printf("s1:%+v, addr:%p, s1 addr:%p\n", s1, s1, &s1)

	s1.update("tom")
	fmt.Printf("s1:%+v, addr:%p\n", s1, s1)
}
