package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

func solve() {
	n := readInt()
	a := make([]int, n)
	for i := range a {
		a[i] = readInt()
	}

	back := 0
	for i := n - 1; i > 0; i-- {
		if a[i-1] < a[i] {
			back = i
			break
		}
	}

	front := -1
	for front < back && len(a[front+1:]) > 0 {
		a = a[front+1:]
		back -= front + 1
		front = 0
		for i := 1; i < len(a); i++ {
			if a[i-1] > a[i] {
				break
			}
			front = i
		}
		// fmt.Println(front, back)
	}

	// fmt.Println("done:", front, back)
	// fmt.Println("ans:", n-len(a))
	fmt.Println(n - len(a))
}

func main() {
	k := readInt()
	for i := 0; i < k; i++ {
		solve()
	}
}

var (
	readString func() string
	readBytes  func() []byte
)

func init() {
	log.SetFlags(log.Lshortfile)
	readString, readBytes = newReadString(os.Stdin)
}

func newReadString(ior io.Reader) (func() string, func() []byte) {
	r := bufio.NewScanner(ior)
	r.Buffer(make([]byte, 1024), math.MaxInt32)
	r.Split(bufio.ScanWords)

	f1 := func() string {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Text()
	}
	f2 := func() []byte {
		if !r.Scan() {
			panic("Scan failed")
		}
		return r.Bytes()
	}
	return f1, f2
}

func readInt64() int64 {
	i, err := strconv.ParseInt(readString(), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func readInt() int {
	return int(readInt64())
}

func readFloat64() float64 {
	f, err := strconv.ParseFloat(readString(), 64)
	if err != nil {
		panic(err.Error())
	}
	return f
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sum(a []int) int {
	var ret int
	for i := range a {
		ret += a[i]
	}
	return ret
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func absInt64(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}

func sumInt64(a []int64) int64 {
	var ret int64
	for i := range a {
		ret += a[i]
	}
	return ret
}

func sumFloat64(a []float64) float64 {
	var ret float64
	for i := range a {
		ret += a[i]
	}
	return ret
}

func gcd(m, n int) int {
	for m%n != 0 {
		m, n = n, m%n
	}
	return n
}

func lcm(m, n int) int {
	return m / gcd(m, n) * n
}
