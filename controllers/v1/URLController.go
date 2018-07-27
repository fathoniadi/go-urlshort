package v1

import (
	"net/http"
)

type HomeController struct {
}

func URLDecodeGet(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("welcome"))
	return
}
