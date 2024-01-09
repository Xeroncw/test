package main

import (
	"fmt"
	"encoding/json"
)

type Result struct {
	Code int `json:"code"`
	Message string `json:"msg"`
}

func main() {
	var res Result
	res.Code = 200
	res.Message = "success"
	tojson(&res)

	setData(&res)
	tojson(&res)
}

func tojson(res *Result) {
	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error", errs)
	}
	fmt.Println("json data:", string(jsons))
}

func setData(res *Result) {
	res.Code = 500
	res.Message = "fail"
}