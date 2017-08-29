package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//register root handle for any method
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Print("root get hit")
		w.Header().Set("content-type", "application/json")
		fmt.Fprintf(w, "{\"Hello\":\"Binarian\"}")
	})
	log.Print("server alive")
	log.Fatal(http.ListenAndServe(":8080", nil)) //server will listen to port 8080
}
