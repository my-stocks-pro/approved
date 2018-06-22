package redis

import (
	"net/http"
	"bytes"
)

func Post(data []byte) (*http.Response, error) {

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
