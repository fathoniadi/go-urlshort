package v1

import (
	"github.com/speps/go-hashids"
	"net/http"
)

type HomeController struct {
}

func URLDecodeGet(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("welcome"))
	return
}

func URLEncodeGet(res http.ResponseWriter, req *http.Request) {
	hd := hashids.NewData()
	hd.Salt = "this is my salt"
	hd.MinLength = 5
	h, _ := hashids.NewWithData(hd)
	for i := 1; i <= 10000; i++ {
		e, _ := h.Encode([]int{i})
		res.Write([]byte(e))
		res.Write([]byte("\n"))
	}

	//db.config.findAndModify({query: {parameter : "last"},update:{ $inc: { key: +1, "last_key": 1 } }})
	return
}
