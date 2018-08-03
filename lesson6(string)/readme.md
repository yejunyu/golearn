[toc]
### 字符串长度
```go
s := "神奇大叶子" // UTF-8
fmt.Println(len(s))

// 15
```
熟悉`python`的朋友应该觉得结果就是5,为什么`go`里是15呢,因为一个中文字符占3个字节
我们来遍历一下
```go
for _, b := range []byte(s) {
	fmt.Printf("%X ", b) //  16进制
}
// E7 A5 9E E5 A5 87 E5 A4 A7 E5 8F B6 E5 AD 90

for i, ch := range s {
	fmt.Printf("(%d %X)", i, ch)
}
// (0 795E)(3 5947)(6 5927)(9 53F6)(12 5B50)
```
可见一个中文字符确实是3个字节,那我就想知道我这句话是几个字呢,不想知道字节,变量那章讲过 `rune`就是`go`里的 `char`
`fmt.Println("Rune count:", utf8.RuneCountInString(s))`
### 遍历
```go
bytes := []byte(s)
for len(bytes) > 0 {
	ch, size := utf8.DecodeRune(bytes)
	bytes = bytes[size:]
	fmt.Printf("%c ", ch) // 字符
}
```
这样是不是略显麻烦,通常要遍历字符串,看下面
```go
for i, ch := range []rune(s) {
	fmt.Printf("(%d %c) ", i, ch)
}
// (0 神) (1 奇) (2 大) (3 叶) (4 子)
```
utf8还有很多方法,可以自行尝试
还有 `stings` 包包含了很多字符串操作的方法,千万不要傻兮兮的自己实现

### 总结
- 使用 `range` 遍历pos,rune 对
- 使用 `utf8.RuneCountInString` 获得字符数
- `len` 获取的是字节数
- 使用`[]byte`获得字节
- `strings`包包含很多字符串的操作
    - Fields,Split,Join
    - Contains,Index
    - ToLower,ToUpper
    - Trim,TrimRight,TrimLeft
