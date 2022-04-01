package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func hello(w http.ResponseWriter, req *http.Request) {
	var name, _ = os.Hostname()
	ip := ReadUserIP(req)
	fmt.Printf("ip: %v\n", ip)
	fmt.Fprintf(w, "hello from %v\n", name)
}

func ReadUserIP(r *http.Request) string {
	fmt.Println("headers:", r.Header)
	fmt.Println("X-Real_Ip", r.Header.Get("X-Real-Ip"), "X-Forwarded-For", r.Header.Get("X-Forwarded-For"), "RemoteAddr", r.RemoteAddr)
	IPAddress := r.Header.Get("X-Real-Ip")
	if IPAddress == "" {
		IPAddress = r.Header.Get("X-Forwarded-For")
	}
	if IPAddress == "" {
		IPAddress = r.RemoteAddr
	}
	return IPAddress
}

func main() {
	http.HandleFunc("/", hello)

	port := os.Args[1]
	if _, err := strconv.Atoi(port); err != nil {
		fmt.Printf("Expected first arg to be a port number but receive %v\n", port)
		os.Exit(1)
	}
	fmt.Printf("Heol~, u have a http server running on port %v\n", port)
	http.ListenAndServe(":"+port, nil)
}
