package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

type args struct {
	V bool
}

var input args
var mu sync.Mutex

func init() {
	log.Println("程序初始化Init")

	// 集中参数定义
	input = args{}
	flag.BoolVar(&input.V, "h", false, "this help") // 绑定位置, 命名表示, 默认状态 , 用户提示
	// 默认为 bool 只要用户有输入  就为true
}

func main() {
	// 解析flag
	flag.Parse()
	mu.Lock()
	defer mu.Unlock()
	log.Println(input)
	if input.V {
		//flag.Usage()
		usage()
	}

	log.Println(os.Args)
}

// 帮助文档打印
func usage() {
	fmt.Fprintf(os.Stderr,`this is go flag test 
	Usage: flagTest [-h bool]
	`)
	flag.PrintDefaults()
}