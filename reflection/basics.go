package main

import (
	"fmt"
	"reflect"
	"time"
)

func print(i interface{}, s string) {

	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)
	k := v.Kind()

	fmt.Printf("i: %-12v(%10s)    TypeOf: %-15sValueOf: %-25sKind: %s\n",
		i, s, t.String(), v.String(), k.String())

	switch k {
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fv := v.Field(i)
			ft := fv.Type()

			// The name for the field is not available in the field value.
			// It is only available through the structure
			fName := v.Type().Field(i).Name

			fmt.Printf("      field: %12s    type: %12s   value: %12v\n",
				fName, ft, fv)
		}
	}
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
