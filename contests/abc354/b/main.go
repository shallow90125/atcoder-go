package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var r = bufio.NewScanner(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer flush()

	n := in()
	s := make([]string, n)
	t := 0

	for i := 0; i < n; i++ {
		s[i] = is()
		t += in()
	}

	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	out(s[t%n])
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
