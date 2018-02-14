package main

import (
  "os"
	"fmt"
	"os/exec"
  "log"
  "net/http"
)

func main() {
	port := os.Getenv("port")

	http.HandleFunc("/hugobuild", HugoBuildCommand)
	
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", fs)
	
	log.Println("Starting webserver listening on", port)
	http.ListenAndServe(":"+port, nil)
}

func HugoBuildCommand(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Running Hugo\n")
	
	out, err := exec.Command("hugo").Output()
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Fprintf(w, "%s\n", string(out))
}