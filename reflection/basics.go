package main

import (
	"fmt"
	"reflect"
	"time"
)

func print(i interface{}, s string) {

	ti := reflect.TypeOf(i)
	vi := reflect.ValueOf(i)
	ki := vi.Kind()

	fmt.Printf("i: %-12v(%10s)    TypeOf: %-15sValueOf: %-25sKind: %s\n",
		i, s, ti.String(), vi.String(), ki.String())
}

type Person struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	SSN       uint      `json:"SSN"`
	Birthdate time.Time `json:"Birthdate"`
}

func main() {

	i := 1
	j := uint16(2)
	print(i, "int")
	print(&i, "&int")
	print(j, "uint16")
	print(&j, "&uint16")

	s := "this"
	print(s, "string")
	print(&s, "&string")

	person := Person{
		FirstName: "Bob",
		LastName:  "Barker",
		SSN:       uint(123456789),
		Birthdate: time.Date(1960, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	print(person, "Person")
	print(&person, "&Person")
}
