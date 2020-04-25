package main

import (
	"fmt"
	"reflect"
	"time"
)

func f(arr [5]int) (sl []int){
	sl = arr[:]
	return
}

func main() {
	arr := [5]int{1,2,3,4,5}
	rst := f(arr)
	fmt.Println(rst)
	fmt.Println(reflect.TypeOf(rst))
	time.Sleep(time.Second * 20)
}
