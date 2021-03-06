package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("进入http请求")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
	fmt.Println(req.URL.Path)
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServe("zhengyudeMacBook-Pro.local:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
