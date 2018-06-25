package redis

import (
	"net/http"
	"bytes"
	"fmt"
)

func POST(data []byte) (*http.Response, error) {

	fmt.Println(string(data))

	req, errReq := http.NewRequest("POST", "http://127.0.0.1:8008/data/redis/approved", bytes.NewReader(data))
	if errReq != nil {
		return nil, errReq
	}

	client := http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, errResp
	}

	return resp, nil
}
