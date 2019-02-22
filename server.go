package main

import (
  "os"
	"fmt"
	"time"
	"os/exec"
  "log"
  "net/http"
)

func main() {
	port := os.Getenv("port")

	http.HandleFunc("/hugobuild", HugoBuildCommand)
	http.HandleFunc("/ghpagespush", GHPagesPushCommand)
	
	fs := http.FileServer(http.Dir("public/"))
	http.Handle("/", fs)
	
	log.Println("Starting webserver listening on", port)
	http.ListenAndServe(":"+port, nil)
}

func HugoBuildCommand(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Running Hugo\n")
	
	out, err := exec.Command("hugo", "--config", "local-config.yaml").Output()
	
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Fprintf(w, "%s\n", string(out))
}

func GHPagesPushCommand(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pushing GH-Pages branch\n")
	
	currentTime := time.Now().Format(time.UnixDate)
	
	fmt.Fprintf(w, currentTime)
	
	out, err := exec.Command("./ghpagespush.sh", currentTime).Output()
	if err != nil { log.Fatal(err) }
	
	fmt.Fprintf(w, "%s\n", string(out))
}