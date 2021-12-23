package main

import (
	"fmt"
	"unsafe"
)

type Element struct {
	value interface{}
}

func main() {
	// run the below line to use bcrypt to store and check for password.

	//testcrpyt()
	//testhmac()
	//testJwt()

	fmt.Println(Element{1})
	fmt.Println(unsafe.Sizeof(Element{}))
}
