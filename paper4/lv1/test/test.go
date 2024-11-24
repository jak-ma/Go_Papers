package test

import (
	"fmt"
)

func Delete() {
	var t int
	fmt.Scan(&t)
	for i := 1; i <= t; i++ {
		var n int
		fmt.Scan(&n)
		m := make(map[int]bool, n)
		mi := make([]int, n)
		for j := 0; j < n; j++ {
			var ni int
			fmt.Scan(&ni)
			mi[j] = ni
			m[mi[j]] = true
		}
		for j := 0; j < n; j++ {
			if m[mi[j]] {
				fmt.Printf("%d ", mi[j])
				m[mi[j]] = false
			}
		}
		fmt.Println()
	}
}
