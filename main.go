package main

import (
	"fmt"
	"log"
	"net/http"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	component := page("")
	component.Render(r.Context(), w)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.Form.Get("username")
	fmt.Println("got username: ", username)

	component := hello(username)
	component.Render(r.Context(), w)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			getHandler(w, r)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - only GET allowed!"))
	})

	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			postHandler(w, r)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - only POST allowed!"))
	})

	// Start the server.
	fmt.Println("listening on http://localhost:8000")
	if err := http.ListenAndServe("localhost:8000", mux); err != nil {
		log.Printf("error listening: %v", err)
	}
}
