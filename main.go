package main

import "fmt"

// IsEven kiểm tra xem số nguyên n có phải là số chẵn không.

func main() {
	num := 10
	if IsEven(num) {
		fmt.Printf("%d là số chẵn\n", num)
	} else {
		fmt.Printf("%d không phải là số chẵn\n", num)
	}
}
