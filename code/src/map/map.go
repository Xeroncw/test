package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	res := make(map[string]interface{})
	res["code"] = 200
	res["msg"] = "success"
	res["data"] = map[string]interface{}{
		"username" : "Tom",
		"age" : 30,
		"hobby" : []string{"读书", "爬山"},
	}
	fmt.Println("map data:", res)

	jsons, errs := json.Marshal(res)
	if errs != nil {
		fmt.Println("json marshal error", errs)
	}
	fmt.Println("")
	fmt.Println("--- map to json ---")
	fmt.Println("json data:", string(jsons))

	res2 := make(map[string]interface{})
	errs = json.Unmarshal([]byte(jsons), &res2)
	if errs != nil {
		fmt.Println("json unmarshal error", errs)
	}
	fmt.Println("")
	fmt.Println("--- json to map ---")
	fmt.Println("map data:", res2)
}