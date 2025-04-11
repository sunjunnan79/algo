package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println(isValid("{()}["))
}

func isValid(s string) bool {
	l := list.New()
	l.Init()
	l.PushBack("T") //用于站位
	for _, v := range s {
		switch v {
		case '}':
			e := l.Back()
			value, ok := e.Value.(int32)
			if !ok {
				return false
			}

			if value == '{' {
				l.Remove(e)
			} else {
				return false
			}
		case ']':
			e := l.Back()
			value, ok := e.Value.(int32)
			if !ok {
				return false
			}

			if value == '[' {
				l.Remove(e)
			} else {
				return false
			}
		case ')':
			e := l.Back()
			value, ok := e.Value.(int32)
			if !ok {
				return false
			}

			if value == '(' {
				l.Remove(e)
			} else {
				return false
			}
		default:
			l.PushBack(v)
		}
	}

	//查是否有剩余值
	if l.Len() > 1 {
		return false
	} else {
		return true
	}
}
