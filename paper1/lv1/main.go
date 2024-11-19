package main

import "fmt"

func main() {
	Print9x9()
}

// lv2
func Print9x9() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
