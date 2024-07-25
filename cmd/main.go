package main

import "github.com/xeuxdev/go-blog/internal/server"

func main() {

	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	println("Hello World!!!!")
}
