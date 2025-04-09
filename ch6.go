package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName string, lastName string, age int) Person {
	return Person{FirstName: firstName, LastName: lastName, Age: age}
}

func MakePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{FirstName: firstName, LastName: lastName, Age: age}
}

func Ch6() {
	MakePerson("John", "Doe", 55)
	MakePersonPointer("Jane", "Doe", 50)

	strsUpdate := []string{"foo", "bar"}
	fmt.Println(strsUpdate)
	UpdateSlice(strsUpdate, "baz")
	fmt.Println(strsUpdate)

	strsGrow := []string{"qux", "fux"}
	fmt.Println(strsGrow)
	GrowSlice([]string{"qux", "fux"}, "bux")
	fmt.Println(strsGrow)

	people := make([]Person, 0, 100000000)

	for range 100000000 {
		people = append(people, Person{FirstName: "Foo", LastName: "Bar", Age: 10})
	}

	fmt.Println("len(people):", len(people))
}

func UpdateSlice(strs []string, str string) {
	strs[len(strs)-1] = str

	fmt.Println("UpdateSlice:", strs)
}

func GrowSlice(strs []string, str string) {
	strs = append(strs, str)

	fmt.Println("GrowSlice:", strs)
}
