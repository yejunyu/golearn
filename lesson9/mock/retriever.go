package mock

import "fmt"

type Retriever struct {
	Contents string
}

func (r Retriever) String() string {
	return fmt.Sprintf("重写的 tostring\n")
}

/**
语言本身并不需要说明,我继承了 Retriever 这个接口,我只要实现 get 方法就行了
*/
func (r Retriever) Get(url string) string {
	return r.Contents + url
}
