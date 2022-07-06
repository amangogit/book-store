package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// parse json data in request body using Unmarshal
func ParseBody(req *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(req.Body); err != nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
