package main

import (
	"errors"
	"fmt"
)

func f(s string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("出现异常，异常为: %v\n", err)
		}
	}()
	if s == "p" {
		panic(errors.New("自定义错误"))
	}
	fmt.Println(s)
}

func main() {
	f("11111")
	f("p")
	f("22222")
}