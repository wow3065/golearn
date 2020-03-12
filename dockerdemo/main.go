package main

import (
	"fmt"
	"net/http"
	"os"
)

func getHostName(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Print("get hostname failed")
	}

	fmt.Fprintf(w, "%s\n", hostname)
}

func main() {
	http.HandleFunc("/gethostname", getHostName)
	http.ListenAndServe(":18080", nil)
}
