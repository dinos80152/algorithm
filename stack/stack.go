// 1. integer stack, pop, push;
// 2. min;
// 3. o(1) for pop, push, min;
// 4. thread safety
// 5. non-blocking

// Alibaba 研发效能事业部-高级JAVA开发工程师/专家-北京/杭州 interview

package stack

import "fmt"

// {1,3,9,7,5}
// {5,7,9,3,1}
// {5,7,9,3,1,2}
func main() {
	s := Stack{
		minTemp: make(map[int]int),
	}
	s.Push(5)
	fmt.Println(s)
	s.Push(4)
	fmt.Println(s)
	s.Push(3)
	fmt.Println(s)
	s.Push(6)
	fmt.Println(s)
	s.Push(7)
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Pop()
	fmt.Println(s)
	s.Push(4)
	fmt.Println(s)
	s.Push(2)
	fmt.Println(s)
}

type Stack struct {
	elements []int
	length   int
	min      int
	minTemp  map[int]int
}

func (s *Stack) Push(element int) {
	s.elements = append(s.elements, element)
	s.length++
	if s.length == 1 || element < s.min {
		s.min = element
	}
	s.minTemp[s.length] = s.min
}

func (s *Stack) Pop() int {
	last := s.elements[s.length-1]
	s.elements = s.elements[:s.length-1]
	delete(s.minTemp, s.length)
	s.length--
	s.min = s.minTemp[s.length]
	return last
}

func (s *Stack) Min() int {
	return s.minTemp[s.length]
}
