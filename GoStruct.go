// Go Struct is on the line
package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person
	pers1.name = "John"
	pers1.age = 45
	pers1.job = "Developer"
	pers1.salary = 1000
	pers2.name = "Mary"
	pers2.age = 25
	pers2.job = "Designer"
	pers2.salary = 2000
	fmt.Println("name:", pers1.name)
	fmt.Println("age:", pers1.age)
	fmt.Println("Job:", pers1.job)
	fmt.Println("Salary:", pers1.salary)
	fmt.Println("Second person data")
	fmt.Println("name:", pers2.name)
	fmt.Println("age:", pers2.age)
	fmt.Println("Job:", pers2.job)
	fmt.Println("Salary:", pers2.salary)
	
}
