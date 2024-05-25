package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var r = bufio.NewScanner(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer flush()

	n, t := in2()

	m := make(map[int]int)

	for i := 0; i < t; i++ {
		a := in()

		for j := 0; j < n; j++ {
			if (a-1)%n == j {
				m[j]++
				if m[j] == n {
					out(i + 1)
					return
				}
			}

			if (a-1)/n == j {
				m[j+n]++
				if m[j+n] == n {
					out(i + 1)
					return
				}
			}
		}

		if (a-1)%n == (a-1)/n {
			m[n*2]++
			if m[n*2] == n {
				out(i + 1)
				return
			}
		}

		if (a-1)/n == n-(a-1)%n-1 {
			m[n*2+1]++
			if m[n*2+1] == n {
				out(i + 1)
				return
			}
		}
	}

	out(-1)
}

const inf = math.MaxInt64

func init() {
	r.Buffer([]byte{}, math.MaxInt64)
	r.Split(bufio.ScanWords)
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := os.ReadFile("./input.txt")
		if e != nil {
			panic(e)
		}
		r = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}
}

func flush() {
	e := w.Flush()
	if e != nil {
		panic(e)
	}
}

func out(v ...interface{}) {
	_, e := fmt.Fprintln(w, v...)
	if e != nil {
		panic(e)
	}
}

func in() int {
	r.Scan()
	i, e := strconv.Atoi(r.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func in2() (int, int) {
	return in(), in()
}

func in3() (int, int, int) {
	return in(), in(), in()
}

func inn(n int) []int {
	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = in()
	}
	return v
}

func ib() []byte {
	r.Scan()
	return r.Bytes()
}

func ibn(n int) [][]byte {
	v := make([][]byte, n)
	for i := 0; i < n; i++ {
		v[i] = ib()
	}
	return v
}

func is() string {
	r.Scan()
	return r.Text()
}

func isn(n int) []string {
	v := make([]string, n)
	for i := 0; i < n; i++ {
		v[i] = is()
	}
	return v
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func abs(a int) int {
	if 0 < a {
		return a
	}
	return -a
}
