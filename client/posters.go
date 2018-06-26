package client

import (
	"net/http"
	"bytes"
)

func (a *ApprovedType) POST(data []byte) (*http.Response, error) {

	req, errReq := http.NewRequest("POST", a.Config.ApiRedisPostURL, bytes.NewReader(data))
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
