package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	var an []int

	for _, input := range inputs {
		a, _ := strconv.Atoi(input)
		an = append(an, a)
	}

	for i := 0; i < len(an)-1; i++ {
		fmt.Print(an[i] * an[i+1])
		if i < len(an)-2 {
			fmt.Print(" ")
		}
	}
}
