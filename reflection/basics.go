package main

import (
	"fmt"
	"reflect"
)

func print(i interface{}, s string) {

	ti := reflect.TypeOf(i)
	vi := reflect.ValueOf(i)
	ki := vi.Kind()

	fmt.Printf("i: %-12v(%10s) TypeOf: %-12sValueOf: %-17sKind: %s\n",
		i, s, ti.String(), vi.String(), ki.String())
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
}
