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

	n := in()

	d := make(map[int]int)

	for i := 0; i < n; i++ {
		a, c := in2()

		if d[c] == 0 {
			d[c] = a
		} else if a < d[c] {
			d[c] = a
		}
	}

	o := 0

	for _, v := range d {
		if o < v {
			o = v
		}
	}

	out(o)
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
