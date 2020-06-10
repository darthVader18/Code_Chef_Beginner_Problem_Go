package main

import "fmt"

func main(){
	var t int
	fmt.Scan(&t)
	for index := 0; index < t; index++ {
		var n int
		fmt.Scan(&n)
		fmt.Println(n/2 +1)
	}
}
