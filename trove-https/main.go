package main

import (
	"flag"
	"log"
	"net/http"
)

type flags struct {
	path string
}

func main() {

	f, err := getFlags()
	if err != nil {
		log.Fatalf("flags parsing fail: %v", err)
	}

	fs := http.FileServer(http.Dir(f.path))
	http.Handle("/", http.StripPrefix("/", fs))

	go func() {
		err := http.ListenAndServe(":80", nil)
		if err != nil {
			log.Fatalf("ListenAndServe: ", err)
		}
	}()

	err = http.ListenAndServeTLS(":443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatalf("ListenAndServeTLS: ", err)
	}
}

func getFlags() (flags, error) {

	p := flag.String("path", "/trove_files", "file folder")

	flag.Parse()

	return flags{*p}, nil
}
