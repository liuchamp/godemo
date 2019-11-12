package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	pa()
}

func randtest() {
	// 例如`rand.Intn`返回一个整型随机数n，0<=n<100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// `rand.Float64` 返回一个`float64` `f`,
	// `0.0 <= f < 1.0`
	fmt.Println(rand.Float64())

	// 这个方法可以用来生成其他数值范围内的随机数，
	// 例如`5.0 <= f < 10.0`
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// 为了使随机数生成器具有确定性，可以给它一个seed
	s1 := rand.NewSource(42)
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	// 如果源使用一个和上面相同的seed，将生成一样的随机数
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()
}

// 如果想获取比较真的随机数， 可以将随机seed 换成如UnixNano

func randtest2() {
	fmt.Println(getRandInt(56))
}

func getRandInt(rn int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	if rn <= 0 {
		n := r.Intn(100)
		return n
	} else {
		return r.Intn(rn)
	}

}

//解析十六进制
func pa() {
	s := strings.Replace("0x86", "0x", "", 1)
	c, ok := strconv.ParseInt(s, 16, 0)
	if ok == nil {
		fmt.Print(c)
	}
}
