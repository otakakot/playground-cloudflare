package function

import "net/http"

func Hello(w http.ResponseWriter, req *http.Request) {
	msg := "Hello!"
	w.Write([]byte(msg))
}
