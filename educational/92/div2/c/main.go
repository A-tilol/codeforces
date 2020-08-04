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

var numbers = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func solve() {
	b := readBytes()
	ans := math.MaxInt32
	for _, x := range numbers {
		for _, y := range numbers {
			removed := make([]byte, 0)
			for i := range b {
				if b[i] == x || b[i] == y {
					removed = append(removed, b[i])
				}
			}

			if x == y {
				ans = min(ans, len(b)-len(removed))
				continue
			}

			cnt := 0
			pre := byte(' ')
			for _, x := range removed {
				if x != pre {
					cnt++
					pre = x
				}
			}
			// fmt.Println(x, y, len(b), len(removed), cnt)
			if cnt%2 == 1 {
				cnt--
			}
			ans = min(ans, len(b)-cnt)
		}
	}
	fmt.Println(ans)
}

func main() {
	k := readInt()
	for i := 0; i < k; i++ {
		solve()
	}
}

// sort ------------------------------------------------------------

type xxx struct {
	x int
}

type sortArray []xxx

func (s sortArray) Len() int           { return len(s) }
func (s sortArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortArray) Less(i, j int) bool { return s[i].x < s[j].x }

// -----------------------------------------------------------------

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

func readInt() int {
	return int(readInt64())
}

func readInt64() int64 {
	i, err := strconv.ParseInt(readString(), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	return i
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

func maxInt64(a, b int64) int64 {
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
