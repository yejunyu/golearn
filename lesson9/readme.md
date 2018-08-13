#### Table of contents
- [duck typing和面向接口编程](#duck-typing和面向接口编程)
- [总结](#总结)

### duck typing和面向接口编程
> 百度百科的解释
https://baike.baidu.com/item/%E9%B8%AD%E5%AD%90%E7%B1%BB%E5%9E%8B/10845665

自行查看百科,简单的说,就是一个长得像鸭子,而且也有鸭子的特点的,我们就可以称他为鸭子
但是每个人(使用者)的理解都不同,孩子可能觉得黄黄的,扁嘴巴就是鸭子,吃货可能觉得要能吃的长翅膀的才叫鸭子,简而言之,是不是鸭子是由使用者决定的,重点在于这个对象(鸭子)他能提供什么功能

说回代码,我现在想实现一个`download`功能,功能就是 get 一个 url 的内容
动态语言`python`是如何实现的呢
```
def download(retriever):
    return retriever.get(url)
```
动态语言实现鸭子类型很方便,但是有两个小缺点
1. 运行时才知道传入的 retriever 有没有 get 方法
2. 需要注释来说明接口(download 需要传入一个有 get 方法的对象)

那静态语言`java`是怎么实现鸭子类型的呢,`java`其实没有鸭子类型,但是有接口的继承
```java
<R extends Retriever>
String dowload(R r){
    return r.get(url)
```
Retriever 里面有 get 方法,就必须要实现Retriever 接口
![image](http://upload-images.jianshu.io/upload_images/5317015-ee9a9378fc0b52ec.jpg?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)

传统的面向对象语言中,接口是实现者定义的,在`java`里,假如我有个`file`接口,里面又`read`和`write`两个方法,相当于告诉了别人,我有这两个方法,你怎么用我不管,你继承我这个接口,自行实现.
而`go`里接口是使用者定义的
比如刚刚的 download 功能
```go
type Retriever interface {
	Get(url string) string
}
func main() {
	var r Retriever
	fmt.Println(download(r))
}
```
使用者定义了接口
具体怎么实现的使用者不用管,与`java`相反,如果想扩展接口,就用上一章学的方法
简而言之,我只关心接口里提供了什么功能(有什么方法)
来看一下, download 这个完整的代码
首先我们有一个 retriever.go 这个包
```go
package mock

import "fmt"

type Retriever struct {
	Contents string
}

/**
语言本身并不需要说明,我继承了 Retriever 这个接口,我只要实现 get 方法就行了
 */
func (r Retriever) Get(url string) string {
	return r.Contents + url
}
```
`main` 包
```go
package main

import (
	"golearn/lesson9/mock"
	"fmt"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	url := "abcd"
	return r.Get(url)
}

func main() {
	var r Retriever
	r = mock.Retriever{"this is a fake mock"}
	fmt.Println(download(r))
        // this is a fake mockabcd
}
```

接口的实现是隐式的,我只需要实现接口的方法就行了
有一个注意的点
`r = mock.Retriever{"this is a fake mock"}`
写成
`r = &mock.Retriever{"this is a fake mock"}`
也就是说,传一个值过去也行,传一个指针过去也行,因为 Get 接收的是一个值,如果接收的是指针,就只能传指针
- 接口变量自带指针
- 接口变量同样采用值传递,几乎不需要使用接口的指针
- 指针接收者只能接收指针,值接收者两者都可以

来学一下查看接口变量的类型
- 表示任何类型: `interface{}`
- Type Assertion
- Type Switch

```go
// type switch
func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents: ", v.Contents)
	case *mock.Retriever:
		fmt.Println("point: ", v.Contents)
}

// Type assertion
if mockRetriever, ok := r.(mock.Retriever); ok {
	fmt.Println(mockRetriever.Contents)
} else {
	fmt.Println("not a mock retriever")
}
```

下面看一个最常用的接口, toString
在`go`里叫 stringer
```go
func (r Retriever) String() string {
	return fmt.Sprintf("重写的 tostring\n")
}
```
因为这里是值接收者,所以可以传值也可以传指针
```go
var r Retriever
r = mock.Retriever{"this is a fake mock"}
fmt.Println(r)
// 重写的 tostring
```

### 总结
- 无须继承也不用关心是什么接口,只需要关心接口提供的功能即可
- 方法如果是值接收者,既可以接收值也可以接收指针,如果是指针接收者,只能接收指针



