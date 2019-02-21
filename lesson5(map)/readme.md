#### Table of contents
- [声明一个 map](#声明一个-map)
- [遍历一个 map](#遍历一个-map)
- [获得一个 value和删除一个 value](#获得一个-value和删除一个-value)
    - [删除一个元素](#删除一个元素)
- [总结](#总结)
### 声明一个 map
```go

m := map[string]string{
		"name":     "神奇大叶子",
		"age":      "26",
		"language": "golang",
	}
m2 := make(map[string]string) // map[]
var m3 map[string]string // map[]
fmt.Println(m, m2, m3)
```
三种声明方式,前文说道的,推荐用 `make`, 这里指出一下, `make` 创建的是空 `map`,`var` 声明的没赋值的是 `nil`，空的map依旧是map，类型是map而不是nil有本质区别
看段代码
```go
fmt.Println(m2 == nil, m3 == nil)
// false    true
// 但是这里打印出来都是map[]
```
两者是有区别的,一定要注意
下面讲讲遍历
### 遍历一个 map
```go
for key, value := range m {
		fmt.Println(key + " => " + value)
	}

/**
name => 神奇大叶子
age => 26
language => golang
*/
```
### 获得一个 value和删除一个 value
和`java`一样,`go`内部的 `key` 也是通过 hash 得到的
要想得到一个 `value`很简单
```go
name:= m["name"]

// 神奇大叶子
```
这里有个小问题,要是我`key`写错了怎么办,通常其他语言会报错,找不到 `key`对吧,看看`go`里是怎么样的
```
name1:= m["name1"]
fmt.Println(name1)

```
返回了一个空,不是 `nil`, 而是一个空
那怎么判断 `map`有没有这个 `key`呢,和前面的 `file`读文件一样
```go
if name1, ok := m["name1"]; ok {
	fmt.Println(name1)
} else {
	fmt.Println("Key not exist")
}
```
`map`是有两个返回值的,一个是 `value`,一个是一个 `bool`类型,`true`代表有这个`key`
#### 删除一个元素
```go
fmt.Println("Deleting values")
delete(m, "age")
age, ok := m["age"]
fmt.Println(age, ok)

// Deleting values
//  false
```
### 总结
- `map` 使用`哈希表`,`key`必须是可以比较相等的类型比如`int`,`string`
- 除了 `slice,map,function` 的内建类型,其余的类型都可以作为 `key`, 这是语言内部实现的,不用像 `java` 一样自己重写 `hash` 方法
- `Struct` 类型里面不包含上述字段的,也可以作为`key`

