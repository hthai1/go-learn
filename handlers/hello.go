package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)
	h.l.Printf("Data %s\n", d)
	//io.WriteString(w, "Hello !\n")
	fmt.Fprintf(rw, "Your data %s\n", d)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
}
