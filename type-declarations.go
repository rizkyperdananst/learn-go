package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpRizky NoKTP = "1231543456341"
	var marriedStatus Married = true

	fmt.Println(noKtpRizky)
	fmt.Println(marriedStatus)

	// byte()
}