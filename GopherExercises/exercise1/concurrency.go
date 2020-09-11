package main

import (
	"fmt"
	"time"
)

func main() {
	go counter("abc")
	counter("hello")
}
func counter(str string) {
	i := 0
	for i < 20 {
		fmt.Println(str)
		time.Sleep(time.Millisecond * 300)
		i = i + 1
	}
}
