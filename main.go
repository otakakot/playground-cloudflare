package main

import (
	"net/http"

	"github.com/syumai/workers"

	"github.com/otakakot/playground-cloudflare/internal/function"
)

func main() {
	http.HandleFunc("/health", function.Health)

	http.HandleFunc("/echo", function.Echo)

	workers.Serve(nil) // use http.DefaultServeMux
}
