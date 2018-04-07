package main

import (
	"fmt"
	"spiderProject/retriever/real"
	"spiderProject/retriever/mock"
	"time"
	"io/WriteCloser"
)

type Retriever interface {
	Get(url string) string
} 

func download(r Retriever) string {
	return r.Get("http://www.bing.com")
}

func main() {
	var r Retriever
	r = mock.Retriever{"This is lll"}
	fmt.Printf("%T %v\n", r, r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut: time.Minute,
	}
	fmt.Printf("%T %v\n", r, r)

	//fmt.Println(download(r))


}
