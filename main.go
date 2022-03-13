package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	var iid = os.Getenv("RENDER_INSTANCE_ID")
	r := chi.NewRouter()
	r.Get("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("serving index from %s", iid)
		fmt.Fprintln(rw, index)
	}))
	r.Get("/mystyle.css", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("serving index from %s", iid)
		rw.Header().Set("Content-Type", "text/css")
		fmt.Fprintln(rw, css)
	}))
	log.Fatal(http.ListenAndServe(":10000", r))
}

const index = `
<!DOCTYPE html>
<html>
	<head>
		<title>Deploy test</title>
		<link rel="stylesheet" href="mystyle.css">
	</head>
	<body>
		<h1 class='header'>New</h1>
	</body>
</html>
`

const css = `
.header {
	color: green;
}
`
