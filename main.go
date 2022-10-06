package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	hostPort := "127.0.0.1:8080" // default host and port
	newHostPort := ""
	path := os.Args[1]

	fmt.Printf("Input custom host and port or press enter to use default (%v): ", hostPort)

	_, err := fmt.Scanln(&newHostPort)

	if err == nil {
		// input custom hostPort
		hostPort = newHostPort
	}

	fmt.Printf("Server root path: %v\n", path)
	fmt.Printf("Server url: http://%v\n", hostPort)
	panic(http.ListenAndServe(fmt.Sprintf("%v", hostPort), http.FileServer(http.Dir(path))))
}
