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

	na := in()
	a := make([]int, na)
	for i := 0; i < na; i++ {
		a[i] = in()
	}

	nb := in()
	b := make([]int, nb)
	for i := 0; i < nb; i++ {
		b[i] = in()
	}

	nc := in()
	c := make([]int, nc)
	for i := 0; i < nc; i++ {
		c[i] = in()
	}

	nx := in()
	x := make([]int, nx)
	for i := 0; i < nx; i++ {
		x[i] = in()
	}

	m := map[int]bool{}

	for _, va := range a {
		for _, vb := range b {
			for _, vc := range c {
				m[va+vb+vc] = true
			}
		}
	}

	for _, vx := range x {
		if m[vx] {
			out("Yes")
		} else {
			out("No")
		}
	}
}

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

func in4() (int, int, int, int) {
	return in(), in(), in(), in()
}

func ib() []byte {
	r.Scan()
	return r.Bytes()
}

func is() string {
	r.Scan()
	return r.Text()
}
