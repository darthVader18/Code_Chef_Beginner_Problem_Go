package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
/*
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numOfInput, numConvertErr := strconv.Atoi(scanner.Text())
	if numConvertErr != nil {
		panic(numConvertErr)
	}

	for i := 0; i < numOfInput; i++ {
		scanner.Scan()
		num := ""
		numString := scanner.Text()
		for j := len(numString) - 1; j >= 0; j-- {
			num += string(numString[j])
		}
		finalNum, err := strconv.Atoi(num)
		if err != nil {
			panic(err)
		}
		fmt.Println(finalNum)
	}
}
 */

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	Reverse(t, scanner)
}

func Reverse(t int, scanner *bufio.Scanner) {
	for index := 0; index < t; index++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		result := 0
		for n > 0 {
			rem := n % 10
			result = (result * 10) + (rem)
			n = n / 10
		}
		fmt.Println(result)
	}
}
