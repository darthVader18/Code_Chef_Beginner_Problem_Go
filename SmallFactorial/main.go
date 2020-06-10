package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for index := 0; index < t; index++ {
		var n, fact int
		fact = 1
		fmt.Scan(&n)
		for i := 1; i <= n; i++ {
			fact = fact * i
		}
		fmt.Println(fact)
	}
}
