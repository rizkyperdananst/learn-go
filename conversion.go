package main

import "fmt"

func main() {
	var nilai32 int32 = 129
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Rizky"
	var e byte = name[0] // fungsinya untuk membuat string dalam variabel name ini diubah menjadi byte
	var eString string = string(e) // mengubah byte ke string

	fmt.Println(name)
	fmt.Println(eString)
}