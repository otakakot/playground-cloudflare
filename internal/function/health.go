package function

import (
	"net/http"
)

func Health(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("OK"))
}
