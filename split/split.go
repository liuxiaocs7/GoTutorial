package split

import (
	"fmt"
	"strings"
)

// split.go

// Split 将s按照sep进行分割,返回要给字符串的切片
// Split("我爱你", "爱") => ["我", "你"]
func Split(s, sep string) (ret []string) {
	ret = make([]string, 0, strings.Count(s, sep)+1)
	idx := strings.Index(s, sep)
	for idx > -1 {
		ret = append(ret, s[:idx])
		s = s[idx+len(sep):]
		idx = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}

func main() {
	ret := Split("我爱你", "爱")
	fmt.Println(ret)
}
