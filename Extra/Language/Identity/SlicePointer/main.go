package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := []byte{0, 1, 2, 3}
	c := a[1:]
	d := a[1:]

	fmt.Println(reflect.ValueOf(c).Pointer() == reflect.ValueOf(d).Pointer())
	fmt.Println(reflect.ValueOf(a).Pointer() == reflect.ValueOf(d).Pointer())

	ahdr := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	chdr := (*reflect.SliceHeader)(unsafe.Pointer(&c))

	fmt.Printf("%+v\n", ahdr)
	fmt.Printf("%+v\n", chdr.Data-1)
	chdr.Data = chdr.Data - 1
	chdr.Len = chdr.Len + 1

	fmt.Println("a", a)
	fmt.Println("c", c)
}
