package project_2

import (
	"fmt"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func Prac() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	x := 10
	if x < 5 {
		fmt.Println("x is less than 5")
	} else {
		fmt.Println("x is not less than 5")
	}
	day := "Wednesday"
	switch day {
	case "Monday":
		fmt.Println("It's Monday")
	case "Wednesday":
		fmt.Println("It's Wednesday")
	default:
		fmt.Println("It's some other day")
	}
	defer fmt.Println("This will be executed last")
	fmt.Println("This will be executed first")
	a := 42
	p := &a // p is a pointer to x
	fmt.Println(*p)
	person := Person{
		FirstName: "John",
		LastName:  "Doe",
		Age:       30,
	}
	fmt.Println(person.FirstName, person.LastName, "is", person.Age, "years old")
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers[0])
}
