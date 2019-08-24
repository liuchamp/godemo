# Go Demo
学习和使用go。

## [指针](https://github.com/liuchamp/godemo/tree/feature/point)
go语言中指针，采用分离的方式，同时去掉了c/c++中的偏移，避免出现野指针的情况。导致程序异常崩溃。
## strings库

## [函数](https://github.com/liuchamp/godemo/tree/feature/func)
在go语言中，需要注意匿名函数访问外部变量。函数也可以作为变量来使用，实现会回调等。

通过函数变量，可以实现类似于netty中的[链式调用](https://github.com/liuchamp/godemo/commit/6a8b61eb94319e46d38cfa18aa2713b1a83e1cf2)。

[闭包函数的记忆特性](https://github.com/liuchamp/godemo/commit/cfcda04a2bb1610c419489b3c95ac8cfefbb781b)，在开发中经常使用。
## redis学习
本次学习是使用的redis官方库. 
1. [单redis使用](https://github.com/liuchamp/godemo/tree/feature/redis/single)
2. [redis集群使用](https://github.com/liuchamp/godemo/tree/feature/redis/cluster)

## MongoDB学习
### [mgo](https://gopkg.in/mgo.v2)
1. [单客户端连接]()
2. [集群连接]()

开发中，客户端的session model的处理。
