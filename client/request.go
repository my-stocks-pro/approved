package client

import (
	"net/http"
	"fmt"
	"io/ioutil"
)


func (a *ApprovedType) NewRequest() ([]byte, error) {

	fmt.Println(a.Config.NewURL)

	req, err := http.NewRequest("GET", a.Config.NewURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Config.Token))
	cookie := http.Cookie{Name: "session", Value: a.Config.Session}
	req.AddCookie(&cookie)

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
