package main

import (
	"fmt"
)

func main() {
	const s = "wbwbwwbwbwbw"
	var w, b int
	fmt.Scanf("%d %d", &w, &b)
	for i := 0; i < len(s); i++ {
		sw, sb := 0, 0
		for j := 0; j < w+b; j++ {
			if s[(i+j)%len(s)] == 'w' {
				sw++
			} else {
				sb++
			}
		}
		if sw == w && sb == b {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
