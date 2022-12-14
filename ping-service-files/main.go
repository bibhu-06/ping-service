package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {

	urls:=strings.Split(os.Getenv("PING_URLS"),",")
	
	for _,url := range urls {
		
		go pingURL(url)

	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting Down")

	//log.Println(urls)	
}


func pingURL(url string) {

	url = strings.TrimSpace(url)

	for {
		_,err := http.Get(url)

		log.Println("Pinging : " + url)

		if err != nil {
			log.Println("There was no error : " + url)
		}

		time.Sleep(5 * time.Second)
	}
}