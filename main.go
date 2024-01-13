package main

import (
	"net/http"

	"github.com/otakakot/playground-cloudflare/internal/function"
	"github.com/syumai/workers"
)

func main() {
	http.HandleFunc("/hello", function.Hello)

	http.HandleFunc("/echo", function.Echo)

	workers.Serve(nil) // use http.DefaultServeMux
}
