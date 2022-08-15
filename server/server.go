package server

import (
	"fmt"
	"log"
	"net/http"
)

func CreateHttpServer() {
	// URI Handling
	http.HandleFunc("/", URIRoute)

	// Start the server
	if err := http.ListenAndServe(":6731", nil); err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
}
