package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	http.HandleFunc("/", httpmirror)
	log.Fatal(http.ListenAndServe(":"+port(), nil))
}

func httpmirror(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mirror mirror on the wall.\nDid you hear my browser call?\n\n"))
	var dump, _ = httputil.DumpRequest(r, true)

	write(dump)
	w.Write(dump)
}

func port() string {
	port := flag.String("port", "3000", "Enter port for web server")
	flag.Parse()
	fmt.Println("Using port: ", *port)

	return *port
}

func write(s []byte) {

	f, err1 := os.OpenFile("httpmirror.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	defer f.Close()

	fmt.Println(string(s))

	_, err2 := f.Write(s)
	if err2 != nil {
		log.Panic(err2)
	}

	f.Sync()
}
