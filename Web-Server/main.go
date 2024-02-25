package main

import (
	"fmt"
	"log"
	"net/http"
)

const P0RT = 8080
func main() {

	fileServer :=  http.FileServer(http.Dir("./static")) // C:/Users/ADMIN/Documents/PERSONAL/GOLANG/Tutorial/GO_PROJECTS/Web-Server/static

	http.Handle("/", fileServer)

	http.HandleFunc("/form", formHandler)

	http.HandleFunc("/hello", helloHandler)

	fmt.Printf ("Server started at port %d", P0RT)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}


}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "hello using fprintf")

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello World!"))
	w.WriteHeader(http.StatusOK)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	for key, value := range r.Form {
		fmt.Fprintf(w, "key: %v, value: %v\n", key, value)
	}

	fmt.Fprintf(w, "POST request successful\n")

	name := r.FormValue("name")

	address := r.FormValue("address")

	fmt.Fprintf(w, "Name = %s\n", name)

	fmt.Fprintf(w, "Address = %s\n", address)
}
