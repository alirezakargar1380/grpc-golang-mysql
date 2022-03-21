package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/alirezakargar1380/grpc_mysql/user_log"
)

func ParseBody(r *user_log.Log, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}
