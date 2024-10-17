package main

import "fmt"

func main() {
	var names [3]string

	// array ini sifatnya static, jadi jika 5 data di arraynya gak bisa ditambahkan, kecuali pakai slice
	names[0] = "Rizky"
	names[1] = "Perdana"
	names[2] = "Nst"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// dibawah ini cara bikin array langsung
	var values = [5]int{
		1,
		2,
		3,
		4,
		5,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values[3])
	fmt.Println(values[4])

	fmt.Print("Jumlah array names : ")
	fmt.Println(len(names))
	fmt.Print("Jumlah array values : ")
	fmt.Println(len(values))
}