package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scan(&t)
	for index := 0; index < t; index++ {
		var a [3]int
		fmt.Scan(&a[0], &a[1], &a[2])
		sort.Ints(a[:])
		fmt.Println(a[1])
	}
}