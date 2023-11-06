package project_3

import (
	"fmt"
)

func Slice() {
	var numbers [5]int = [5]int{11, 22, 33, 44, 55}
	sliceA := numbers[1:4]
	sliceA = append(sliceA, 99)

	fmt.Println(sliceA)

	var names [5]string = [5]string{"Alice", "Bob", "Charlie", "David", "Eve"}
	fmt.Println(names)

	sliceB := names[0:2]
	sliceC := names[1:3]
	fmt.Println(sliceB, sliceC)

	sliceB[0] = "Frank"
	fmt.Println(sliceB, sliceC)
	fmt.Println(names)

	x := make([]int, 0, 5)
	x = append(x, 6, 7, 8, 9, 10)
	fmt.Println(x)

	studentGrades := map[string]float64{
		"math":    95.2,
		"history": 88.7,
		"science": 73.4,
	}
	studentGrades["english"] = 81.1
	wValue := studentGrades["history"]
	fmt.Println(wValue)

	studentGrades["science"] = 79.3
	fmt.Println(studentGrades)

	delete(studentGrades, "english")
	fmt.Println(studentGrades)

	for subject, grade := range studentGrades {
		fmt.Println(subject, grade)
	}

	if grade, ok := studentGrades["math"]; ok {
		fmt.Println(grade, ok)
	} else {
		fmt.Println("Key not found")
	}

	studentMarks := make(map[string]int)

	studentMarks["Alice"] = 92
	studentMarks["Bob"] = 78
	studentMarks["Charlie"] = 85

	for student, marks := range studentMarks {
		fmt.Println(student, marks)
	}
}
