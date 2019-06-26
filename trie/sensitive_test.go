package trie

import (
	"testing"
	"fmt"
)

// 转换性能测试
func Benchmark_SensitiveTransform(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SensitiveTransform("本土无码公司待办列宁测试GM")
	}
}

// 构造性能测试
func Benchmark_Init(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// 加载屏蔽字树
		e := CreateSensitiveTree()
		if e != nil {
			fmt.Println(e)
		}
	}
}

// 检测性能测试
func Benchmark_SensitiveCheck(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SensitiveCheck("本土无码公司待办列宁测试GM")
	}
}

func init() {
	// 加载屏蔽字树
	e := CreateSensitiveTree()
	if e != nil {
		fmt.Println(e)
	}
}
