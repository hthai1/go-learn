package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Println("Hello World")
		d, err := ioutil.ReadAll(req.Body)
		log.Printf("Data %s\n", d)
		//io.WriteString(w, "Hello !\n")
		fmt.Fprintf(w, "Your data %s\n", d)
		if err != nil {
			http.Error(w, "Oops", http.StatusBadRequest)
			return
			// w.WriteHeader(http.StatusBadRequest)
			// w.Write([]byte("Oops"))
			// return
		}
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, req *http.Request) {
		log.Println("Goodbye World")
		io.WriteString(w, "Good bye\n")
	})

	err := http.ListenAndServe(":9090", nil)
	log.Fatal(err)
}
