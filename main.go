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
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			rw.Header().Add("rl-served-by-instance", iid)
			log.Printf("serving %s from %s", r.URL.Path, iid)
			next.ServeHTTP(rw, r)
		})
	})
	r.Get("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(rw, index)
	}))
	r.Get("/mystyle.css", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
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
		<h1 class='header'>New Demo Again</h1>
	</body>
</html>
`

const css = `
.header {
	color: purple;
}
`
