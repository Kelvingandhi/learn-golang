package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

// Here, defining method to act on object type Student
func (s Student) getAge() int {
	return s.age
}

// Rule of Thumb is when we change value of var inside method, ALWAYS use pointer
// It's never hurt to pass pointer anyways

func (s *Student) setAge(age int) {
	s.age = age
}

// Calculate Avg grade
func (s *Student) getAvgGrade() float32 {

	sum := 0

	for _, v := range s.grades {
		sum += v
	}

	return float32(sum) / float32(len(s.grades))
}

// Find max grade
func (s *Student) getMaxGrade() int {
	max := 0

	for _, v := range s.grades {
		if v > max {
			max = v
		}
	}

	return max
}

func main() {
	s := Student{"Jake", []int{85, 93, 78, 73}, 24}
	fmt.Println(s.getAge())

	s.setAge(26)
	fmt.Println(s.getAge()) // In case of struct, it doesn't matter if we use & operator or pass directly if we use pointer other side

	fmt.Println(s.getAvgGrade())
	fmt.Println(s.getMaxGrade())

}
