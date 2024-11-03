package main

import (
	"learn-go/hello-gin/router"
	"log"
)

func main() {
	router := router.InitRouter()
	err := router.Run()
	if err != nil {
		log.Fatalln("gin run error")
	}
}
