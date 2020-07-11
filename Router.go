package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type PageVars struct {
	Date  string
	Owner string
}

func index(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	v := PageVars{t.Format("01-02-2006 Monday"), "Alex"}
	temp, err := template.ParseFiles("Index.html")
	if err == nil {
		temp.Execute(w, v)
	} else {
		fmt.Println(err)
	}
}
func startServer() *http.Server {
	server := &http.Server{Addr: ":3333"}
	http.HandleFunc("/", index)
	go server.ListenAndServe()
	fmt.Println("returning Server")
	return server
}
func main() {
	fmt.Println("starting up")
	server := startServer()
	defer fmt.Println("shutting down....")
	defer server.Shutdown(context.Background())
	fmt.Println("wait a bit ")
	time.Sleep(10 * time.Second)
}
