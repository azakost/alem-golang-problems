package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func main() {

}

type js struct {
	Children struct {
		Div01 struct {
			Children struct {
				Exam struct {
					Children interface{} `json:"children"`
				} `json:"exam"`
			} `json:"children"`
		} `json:"div-01"`
	} `json:"children"`
}

// Standard parsing API function
func ParseAPI(url string, d interface{}) string {
	got, eget := http.Get(url)
	body, ered := ioutil.ReadAll(got.Body)
	emar := json.Unmarshal(body, d)
	if eget != nil || ered != nil || emar != nil {
		return "Error with parsing " + url
	}
	return "Success"
}

func GrabData(url string) js {
	var list js
	ParseAPI(url, &list)
	return list
}
