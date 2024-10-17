package main

import "fmt"

func main() {
	// dalam array months bisa aja dikasih 12 bisa juga ...
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Print("Jumlah slice1 : ")
	fmt.Println(len(slice1))
	fmt.Print("Capasity slice1 : ")
	fmt.Println(cap(slice1))

	// months[5] = "Diubah"
	// fmt.Println(slice1)

	slice1[0] = "Mei Lagi"
	fmt.Println(months)

	// len(slice) = untuk mendapatkan panjang slice
	// cap(slice) = untuk mendapatkan kapasitas slice
	// append(slice, data) = untuk menambahkan slice baru dengan menambahkan data ke posisi terkahir slice, jika kapasitas sudah penuh maka akan membuat slice baru
	// make(slice, length, capacity) = untuk membuat slice baru
	// copy(destination, source) = untuk copy slice
}