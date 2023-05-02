package main

import (
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Name    string `json:"name"`
	NIM     string `json:"nim"`
	Address string `json:"address"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 PAGE NOT FOUND", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "forms.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err : %v", err)
			return
		}

		fmt.Fprintf(w, "Post from website r.postfrom = %v\n", r.PostForm)
		name := r.FormValue("name")
		nim := r.FormValue("nim")
		address := r.FormValue("address")

		fmt.Fprintf(w, "Nama = %s\n", name)
		fmt.Fprintf(w, "NIM = %s\n", nim)
		fmt.Fprintf(w, "Alamat = %s\n", address)

	default:
		fmt.Fprintf(w, "Only get and post")
	}
}
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":2000", nil))
}
