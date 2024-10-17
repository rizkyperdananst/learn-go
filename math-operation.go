package main

import "fmt"

func main() {
	var result = 10 + 10

	fmt.Println(result)

	var a = 10;
	var b = 10;
	var c  = a * b; // operasi matematika

	var i = 10;
	i += 10; // i = i + 10 // augmented assignment
	fmt.Println(i);
	fmt.Println(c)

	a++ // a = a + 1 // disini maksudnya dia bakalan nambah isi didalam a + 1
	fmt.Println(a);

	var negative = -100;
	var positive = +100;
	fmt.Println(negative)
	fmt.Println(positive)
}