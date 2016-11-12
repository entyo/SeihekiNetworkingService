// web.go
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	// Parse flags
	var portNum = flag.Int("port", 8080, "The numboer of port which is used of by this server.")
	flag.Parse()
	port := strconv.Itoa(*portNum)

	http.HandleFunc("/", hello)
	fmt.Printf("Listening at :%d\n", *portNum)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello, world")
	log.Println("served")
}
