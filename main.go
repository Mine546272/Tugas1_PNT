package main

import (
	"io"
	"net/http"
)

func endpoint1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Gaboleh Pinjol ya teman teman")
}

func endpoint2(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Gultik 10 rb/porsi")
}

func getroot(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "endpoint 1 atau 2")
}

func main() {

	http.HandleFunc("/", getroot)
	http.HandleFunc("/endpoint1", endpoint1)
	http.HandleFunc("/endpoint2", endpoint2)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}

}
