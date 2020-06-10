package main
import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scan(&t)
	for index := 0; index < t; index++ {
		var n float64
		fmt.Scan(&n)
		result := math.Sqrt(n)
		fmt.Println(int(result))
	}
}


