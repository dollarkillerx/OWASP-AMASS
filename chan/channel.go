package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	chin := make(chan int, 10)

	group := sync.WaitGroup{}

	group.Add(2)

	go func() {
		defer group.Done()
		for i := range chin {
			log.Println(i)
		}

		// range 如果沒有close 就會阻塞
		//cc:
		//	for {
		//		select {
		//		case k, ok := <-chin:
		//			{
		//				if ok {
		//					log.Println(k)
		//				} else {
		//					break cc
		//				}
		//			}
		//		}
		//	}
		log.Println("end")
	}()

	go func() {
		defer group.Done()
		time.Sleep(time.Second)
		for i := 0; i < 100; i++ {
			chin <- i
			//time.Sleep(time.Millisecond * 500)
		}
		close(chin)
	}()

	group.Wait()
	//time.Sleep(time.Second * 3)
}
