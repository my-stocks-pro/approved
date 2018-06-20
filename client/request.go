package client

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/my-stocks-pro/approved/env"
)


func (a *ApprovedType) NewRequest() ([]byte, error) {

	fmt.Println(a.NewURL)

	req, err := http.NewRequest("GET", a.NewURL, nil)
	if err != nil {
		return nil, err
	}

	if (env.FLAG & env.HEADER) != 0 {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
	}

	if (env.FLAG & env.COOKIE) != 0 {
		cookie := http.Cookie{Name: "session", Value: a.Session}
		req.AddCookie(&cookie)
	}

	client := http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		return nil, err
	}

	body, errReadBody := ioutil.ReadAll(resp.Body)
	if errReadBody != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return body, nil
}
