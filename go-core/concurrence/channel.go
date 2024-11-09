package main

import (
	"fmt"
	"time"
)

func go_run1() {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	//time.Sleep(1 * time.Second)
}

func downloading(fileSize int, c chan struct{}) {
	for i := 0; i < int(fileSize); i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Downloading: %d%\n", 100*i/fileSize)
	}
	fmt.Printf("Downloading: %d%\n", 100)
	//	上传完成后，通过信道通知upload函数
	c <- struct{}{}
}

func upload(fileSize uint) {
	c := make(chan struct{})
	go downloading(3, c)
	//等待上传结束
	<-c
	fmt.Println("上传完成")
}

func producer(c chan int, num int) {
	for i := 0; i < num; i++ {
		c <- i
		time.Sleep(1 * time.Second)
	}
	c <- -1
}

func comsumer(c chan int, finished chan struct{}) {
	for task := range c {
		if task == -1 {
			break
		}
		fmt.Println(task)
	}
	finished <- struct{}{}
}

func dataTransferWithChan() {
	//producer和consumer之间的数据通信
	c := make(chan int, 5)
	//主 gorouine和consumer之间的同步
	finishedChannel := make(chan struct{})
	go producer(c, 10)
	go comsumer(c, finishedChannel)
	<-finishedChannel
	fmt.Println("data transfer finished")
}

func main() {
	dataTransferWithChan()
}
