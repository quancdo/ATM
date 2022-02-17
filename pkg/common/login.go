package common

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ParseRequest(req *http.Request, obj interface{}) error {

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(obj)
	if err != nil {
		fmt.Println("error")
		return err
	}

	return nil
}
