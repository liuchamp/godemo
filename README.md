# Go Demo
学习和使用go。

本项目测试go语言基础中的数组和map使用。

## array的初步容量
```go
    s:=[]int{5}
    fmt.Println(cap(s)) //1
    p:=[]int{5,12,6}
    fmt.Println(cap(p)) //3
    s= append(s, 1,6)
    fmt.Println(cap(s)) //4
```
从这段代码中可知，append函数会自动为数组扩容。并且是2^n。
## 删除元素
go语言不能直接删除元素，可以通过切片的方式删除对应的元素。
```go
a := []int{0, 1, 2, 3, 4}
//删除第i个元素
i := 2
a = append(a[:i], a[i+1:]...)
```