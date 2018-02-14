package main

import (
  "os"
  "log"
  "net/http"
)

func main() {
  log.Println("Hello World")

	port := os.Getenv("port")

	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", fs)
	
	log.Println("Starting webserver listening on", port)
	http.ListenAndServe(":"+port, nil)
}
