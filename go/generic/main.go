package main

import "fmt"

/*=========================
	1. 基本语法
	类型参数
	使用方括号 [] 定义类型参数：
=========================
*/

// 泛型函数
func PrintSlice[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

// 泛型结构体
type Stack[T any] struct {
	items []T
}

// 泛型方法
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.items) == 0 {
		var zero T // 泛型的零值
		return zero
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}
