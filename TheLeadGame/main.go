package main

import (
	"fmt"
)

func main() {
	var N, sum1, sum2, diff, winner int
	fmt.Scan(&N)
	for index := 0; index < N; index++ {
		var Si, Ti int
		fmt.Scan(&Si, &Ti)
		sum1 = sum1 + Si
		sum2 = sum2 + Ti

		if sum1 > sum2 {
			if (sum1 - sum2) > diff {
				diff = sum1 - sum2
				winner = 1
			}
			} else {
				if (sum2 - sum1) > diff {
					diff = sum2 - sum1
					winner = 2
				}
			}
		}
		fmt.Println(winner, diff)
}




