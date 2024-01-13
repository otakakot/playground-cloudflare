package function

import (
	"io"
	"net/http"
)

func Echo(w http.ResponseWriter, req *http.Request) {
	io.Copy(w, req.Body)
}
