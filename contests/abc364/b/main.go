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

	h, w := in2()
	sy := in() - 1
	sx := in() - 1
	c := make([][]byte, h)
	for y := 0; y < h; y++ {
		c[y] = ib()
	}
	x := ib()

	for i := 0; i < len(x); i++ {
		cx, cy := sx, sy

		if x[i] == 'L' && 0 < sx {
			cx--
		} else if x[i] == 'R' && sx < w-1 {
			cx++
		} else if x[i] == 'U' && 0 < sy {
			cy--
		} else if x[i] == 'D' && sy < h-1 {
			cy++
		} else {
			continue
		}

		if c[cy][cx] == '.' {
			sx, sy = cx, cy
		}
	}

	out(sy+1, sx+1)
}

const inf = math.MaxInt64
const mod998244353 = 998244353

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

func itoa(v int) string {
	return strconv.Itoa(v)
}

func atoi(v string) int {
	d, e := strconv.Atoi(v)
	if e != nil {
		panic(e)
	}
	return d
}

func abs(a int) int {
	if 0 < a {
		return a
	}
	return -a
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sqrt(i int) int {
	return int(math.Sqrt(float64(i)))
}

func reverse[T any](s []T) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func sorta[T int | byte](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
}

func sortd[T int | byte](s []T) {
	sort.Slice(s, func(i, j int) bool {
		return s[i] > s[j]
	})
}
