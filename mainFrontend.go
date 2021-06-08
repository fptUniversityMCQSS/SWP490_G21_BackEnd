package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//connect frontend to backend
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}
