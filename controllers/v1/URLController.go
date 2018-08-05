package v1

import (
	//"github.com/speps/go-hashids"
	"encoding/json"
	"fmt"
	"github.com/fathoniadi/go-urlshort/core"
	"github.com/fathoniadi/go-urlshort/models"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

type URLRequest struct {
	Url         string `json:"url"`
	Custom_link string `json:"custom_link"`
	Mode        string `json:"mode"`
}

func URLDecodeGet(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("welcome"))

	db := core.DBConnect()

	var configs []models.Config

	if err := db.DB("go_urlshort").C("config").
		Find(nil).Limit(100).All(&configs); err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(configs)

	return
}

func URLEncodePost(res http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	var url_request URLRequest

	rules := govalidator.MapData{
		"mode": []string{"required"},
		"url":  []string{"required", "url"},
	}

	opts := govalidator.Options{
		Request: req,
		Data:    &url_request,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	e := v.Validate()

	if len(e) > 0 {
		error_req := map[string]interface{}{"data": e, "status": 400}
		json.NewEncoder(res).Encode(error_req)
		return
	}

	switch req.FormValue("mode") {
	case "random":
	case "custom_link":
		custom_link := req.FormValue("custom_link")

		if len(custom_link) <= 4 {
			error_req := map[string]interface{}{"data": map[string]string{"custom_link": "Custom link must be greater than 4 character"}, "status": 400}
			json.NewEncoder(res).Encode(error_req)
			return
		}
	default:
		error_req := map[string]interface{}{"data": map[string]string{"mode": "Mode not supported"}, "status": 400}
		json.NewEncoder(res).Encode(error_req)
		return
	}

	// hd := hashids.NewData()
	// hd.Salt = "this is my salt"
	// hd.MinLength = 5
	// h, _ := hashids.NewWithData(hd)
	// for i := 1; i <= 10000; i++ {
	// 	e, _ := h.Encode([]int{i})
	// 	res.Write([]byte(e))
	// 	res.Write([]byte("\n"))
	// }

	//db.config.findAndModify({query: {parameter : "last"},update:{ $inc: { key: +1, "last_key": 1 } }})
	return
}
