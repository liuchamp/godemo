package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func doubleScore(source float32) (score float32) {
	defer func() {
		fmt.Println("ing")
		if score < 1 || score >= 100 {
			//将影响返回值
			score = source
		}
	}()
	score = source * 2
	fmt.Println("end")
	return

	//或者
	//return source * 2
}

func main() {
	a := 1 //line 1
	b := 2 //2
	calc("30", a, b)
	defer calc("1", a, calc("10", a, b)) //3
	a = 0                                //4
	defer calc("2", a, calc("20", a, b)) //5
	b = 1                                //6

	doubleScore(50.0)
	calc("40", a, b)
}
