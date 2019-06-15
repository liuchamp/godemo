package main

import (
	"fmt"
	"time"
)

func main() {
	testUnix()
}
func testUnix() {
	// 使用Unix和UnixNano来分别获取从Unix起始时间
	// 到现在所经过的秒数和微秒数
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// 注意这里没有UnixMillis方法，所以我们需要将
	// 微秒手动除以一个数值来获取毫秒
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// 反过来，你也可以将一个整数秒数或者微秒数转换
	// 为对应的时间
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
func testTime() {
	p := fmt.Println

	// 从获取当前时间开始
	now := time.Now()
	p(now)

	// 你可以提供年，月，日等来创建一个时间。当然时间
	// 总是会和地区联系在一起，也就是时区
	then := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// 你可以获取时间的各个组成部分
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// 输出当天是周几，Monday-Sunday中的一个
	p(then.Weekday())

	// 下面的几个方法判断两个时间的顺序，精确到秒
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Sub方法返回两个时间的间隔(Duration)
	diff := now.Sub(then)
	p(diff)

	// 可以以不同的单位来计算间隔的大小
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// 你可以使用Add方法来为时间增加一个间隔
	// 使用负号表示时间向前推移一个时间间隔
	p(then.Add(diff))
	p(then.Add(-diff))
}
