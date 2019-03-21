package main 	// package 必须是 main

import (
	"fmt"
	"os"
)

// 函数名必须是 main
func main()  {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1])
	}
    os.Exit(0)
}