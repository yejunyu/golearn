
### defer
`go`里有个关键字`defer`
```go
func tryDefer1() {
	defer fmt.Println(1)
	fmt.Println(2)
}

func tryDefer2() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}
```
这两段代码会输出什么呢
tryDefer1: 2 1
tryDefer1: 3 2 1

所以 `defer`有两个特点
- 在函数退出了才执行,甚至前面有 return 也行
- 按照`defer`的顺序从后往前执行(类似栈)

这种特性很方便,可以用在很多资源释放的场合

- Open/Close
- Lock/Unlock
- PrintHeader/PrintFooter


### 异常处理
`go`里没有那么麻烦的 `try` `catch`,`scala`里同样不推崇传统的`java`那种异常捕获,感觉这是个新的思维,还得多学习学习
举个🌰
```go
func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	defer write.Flush()

	f := fib.Fib()
	for i := 0; i < 20; i++ {
		fmt.Fprint(write, strconv.Itoa(f())+"\n")
	}
}
```
我想把 fib 数写进一个文件里,有经验的同学都知道,打开一个文件需要捕获异常,因为这个文件可能不存在,或者写入因为各种权限问题时报错
`go`里很多函数都是两个返回值,第一个是结果,第二个是异常

`go`很方便在于,你打开一个文件,你就顺手 `defer` `close`掉,不用特意包裹在麻烦的`catch`里
用一个`bufio`写效率会高很多,记得要 `flush`到磁盘里
`Fprint`接收一个 `writer`

再来看个异常处理的例子
打开一个文件,都知道文件可能不存在会抛出异常
![](http://oqb4aabpb.bkt.clouddn.com/18-8-11/94248628.jpg)
跟到源码里看到注释说这个错误会是`*PathError`
那就把`*PathError`单独拎出来
```go
func openFile(path string) string {
	file, err := os.Open(path)
	// 对已知的问题的处理
	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Println(pathError.Err)
		} else {
			// 未知问题特殊处理
			fmt.Println("Unkown error", err)
		}
	}
	return file.Name()
}
```
这就是一个比较正常的错误处理了,但是现实工作中,我们需要统一的异常处理,有时候还需要自定义的异常处理

### 统一的异常处理
另创建一个文件夹,做一个服务端,作用就是在网页上显示文件里的内容
```go
package main

import (
	"net/http"
	"os"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		path := request.URL.Path[len("/list/"):]
		file, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		all, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		writer.Write(all)
	})
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
```
这样就启动了一个 server, 我输入`localhost:8888/list/fib.txt`时可以显示我项目中 fib.txt 中的内容,如果输入的路径不对,当然就`panic(err)`了
来看一下
![](http://oqb4aabpb.bkt.clouddn.com/18-8-11/76127553.jpg)
但是如果网址输错了,访问了一个不存在的文件
![](http://oqb4aabpb.bkt.clouddn.com/18-8-11/4449015.jpg)
这个太难看了,稍微改一下
![](http://oqb4aabpb.bkt.clouddn.com/18-8-11/80251839.jpg)
不直接`panic`
http.Error 有三个参数,第一个是 write 也就是你的网页,第二个是出错信息,第三个是 code
随便输错一个再来看
![](http://oqb4aabpb.bkt.clouddn.com/18-8-11/93261495.jpg)
可是这样也不太好,程序的报错不应该暴露给用户,只需要让用户知道 not found 就行了

把整个 handler 函数抽出来
```go
package filelisting

import (
	"net/http"
	"os"
	"io/ioutil"
)

func Handler(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	writer.Write(all)
	return nil
}
```
返回值是 error
定义一个结构体,对应这个handler 函数
```go
type appHandler func(writer http.ResponseWriter, request *http.Request) error
```
下面用到函数式编程的思想
传这个 error 进去,但是返回 http.HandleFunc 需要的函数
有点像`python`的装饰器
```go
func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log.Warn("Error handling request: %s",
				err.Error())
			var code int
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}
```
一个装饰 error 的函数,返回也是一个函数
在函数体内包装 error
整个程序最后是这样的
```go
package main

import (
	"net/http"
	"golearn/lesson11/filelistserver/filelisting"
	"os"
	"github.com/gpmgo/gopm/modules/log"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log.Warn("Error handling request: %s",
				err.Error())
			var code int
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/list/", errWrapper(filelisting.Handler))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
```

![](http://oqb4aabpb.bkt.clouddn.com/18-8-11/55834589.jpg)
最后页面返回就是想要的结果,而不是直接暴露错误给用户


### recover
说道`panic`就要说`recover`

- 仅在 defer 调用中使用
- 获取 panic 的值
- 如果无法处理,可重新 panic
看代码
```go

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred:", err)
		} else {
			panic(r)
		}
	}()
	panic(errors.New("this is a new error"))
}
```
`panic`和`recover`有点像`try` `catch`其实就是`c`语言的`try` `catch`
我的程序出错了(`panic`)但是我不想让他让程序终止,在我可控的范围内处理他(`recover`)
当我发现他是我知道的类型(`error`),我处理他,如果是我意料之外的东西,我还可以继续`panic`

上面的 http 服务器的例子都是系统异常直接抛出了,实际上,预料的异常应该 `recover`,还要自定义一些异常作为给用户看的
![](http://oqb4aabpb.bkt.clouddn.com/18-8-12/55221516.jpg)
![](http://oqb4aabpb.bkt.clouddn.com/18-8-12/86098590.jpg)
在 Handler 这个方法的文件里,实现 userError 这个接口,抛出一个用户异常
详细代码看`github`


### 总结
- 异常处理要用到 `defer`,`panic`,`recover`
- `go`是互联网时代的`c`,`panic`,`recover`都是`c`过来的思想,其实就是`try`,`catch`
- 意料之中的错误用 `err`, 意想不到的问题才 `panic`(尽量不要用)
