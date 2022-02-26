package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func hello(w http.ResponseWriter, req *http.Request) {
	var name, _ = os.Hostname()
    fmt.Fprintf(w, "hello from %v\n", name)
}

func main() {
    http.HandleFunc("/", hello)

    port := os.Args[1]
    if _, err := strconv.Atoi(port); err != nil {
	    fmt.Printf("Expected first arg to be a port number but receive %v\n", port)
	    os.Exit(1)
    }
    fmt.Printf("Heol~, u have a http server running on port %v", port)
    http.ListenAndServe(":" + port, nil)
}
