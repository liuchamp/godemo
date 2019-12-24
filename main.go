package main

import (
	"fmt"
	"regexp"
	s "strings"
)

func main() {
	lmd5()
}

func lmd5() {
	s := "e10adc3949ba59abbe56e057f20f883e"
	fmt.Println(len(s))
}
func test1() {
	s12 := "1231"
	s.Join([]string{s12, "ewrwer"}, ",")
	fmt.Println(s12)
}
func spl() {
	ssr := "{\"token\": \"c3ba4c4f-a3b3-4747-b82f-9e40736a64c8\"}"
	sr := s.Split(ssr, "\"")
	fmt.Println(sr)
}
func test2() {
	var ss []string
	ss = append(ss, "123123")
	ss = append(ss, "sjadingisdajf")
	sm := s.Join(ss, ",")
	fmt.Println(sm)
}

// 这里给fmt.Println起个别名，因为下面我们会多处使用。
var p = fmt.Println

func sum() {
	// 下面是strings包里面提供的一些函数实例。注意这里的函数并不是
	// string对象所拥有的方法，这就是说使用这些字符串操作函数的时候
	// 你必须将字符串对象作为第一个参数传递进去。
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	// 你可以在strings包里面找到更多的函数

	// 这里还有两个字符串操作方法，它们虽然不是strings包里面的函数，
	// 但是还是值得提一下。一个是获取字符串长度，另外一个是从字符串中
	// 获取指定索引的字符
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}

func testRegex() {
	const urlPattern = `((https|http|ftp|rtsp|mms)?://)?` + // 协议
		`(([0-9a-zA-Z]+:)?[0-9a-zA-Z_-]+@)?` + // pwd:user@
		"(" + ipPattern + "|(" + domainPattern + "))" + // IP 或域名
		`(:\d{1,5})?` + // 端口
		`(/+[a-zA-Z0-9][a-zA-Z0-9_.-]*)*/*` + // path
		`(\?([a-zA-Z0-9_-]+(=.*&?)*)*)*` // query

}

func regexpCompile(str string) *regexp.Regexp {
	return regexp.MustCompile("^" + str + "$")
}

func URL(val interface{}) bool {
	return isMatch(url, val)
}

func isMatch(exp *regexp.Regexp, val interface{}) bool {
	switch v := val.(type) {
	case []rune:
		return exp.MatchString(string(v))
	case []byte:
		return exp.Match(v)
	case string:
		return exp.MatchString(v)
	default:
		return false
	}
}
