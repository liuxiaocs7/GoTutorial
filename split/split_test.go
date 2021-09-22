package split

import (
	"reflect"
	"testing"
)

// 测试split
// 输出详细信息
// go test -v
// 跑指定的测试用例
// go test -run=Split/chinese
// 测试覆盖率
// go test -cover
// 输出到文件
// go test -cover -coverprofile=c.out

// func TestSplit(t *testing.T) {
// 	got := Split("我爱你", "爱")
// 	want := []string{"我", "你"}

// 	// got := Split("a:b:c", ":")
// 	// want := []string{"a", "b", "c"}
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("want:%v got:%v", want, got)
// 	}
// }

func TestSplit(t *testing.T) {
	type test struct {
		input string
		sep   string
		want  []string
	}
	tests := map[string]test{
		"simple":     test{input: "我爱你", sep: "爱", want: []string{"我", "你"}},
		"multi sep":  test{input: "a:b:c", sep: ":", want: []string{"a", "b", "c"}},
		"multi sep2": test{input: "abcd", sep: "bc", want: []string{"a", "d"}},
		"chinese":    test{input: "沙河有沙又有河", sep: "沙", want: []string{"", "河有", "又有河"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := Split(tc.input, tc.sep)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("name: %s failed, want:%v got:%v", name, tc.want, got)
			}
		})
	}
}
