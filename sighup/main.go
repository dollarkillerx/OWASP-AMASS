package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	ic := 0
	ko:
	c := make(chan os.Signal)

	// 監聽消息
	signal.Notify(c)
	log.Println("啟動")

	s := <-c
	log.Println("收到推出信號:", s)
	ic ++
	if ic <3 {
		log.Println("守護進程啟動:", ic)
		goto ko
	}else {
		log.Println("超過最大數量 停止程序")
	}
}
