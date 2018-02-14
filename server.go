package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	prod = flag.Bool("prod", false, "Run in production, don't watch frontend")
	addr = flag.String("addr", ":4000", "address to run http server")
)

func main() {
	if *prod {
		buildFrontend()
	} else {
		go watchFrontend()
	}

	static := http.FileServer(http.Dir("frontend-dist"))
	http.Handle("/dist/", http.StripPrefix("/dist/", static))
	http.Handle("/", http.HandlerFunc(serveIndex))

	log.Println("Server running at " + *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Panic(err)
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "frontend-dist/index.html")
}
