package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)





func Req(newIDs []string) {


	url := "https://api.shutterstock.com/v2/images?" + q

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	client := http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		fmt.Println(errResp)
	}

	body, errReadBody := ioutil.ReadAll(resp.Body)
	if errReadBody != nil {
		fmt.Println(errReadBody)
	}

	fmt.Println(string(body))

	tmp := ImageType{}

	errUnm := json.Unmarshal(body, &tmp)
	if errUnm != nil {
		fmt.Println(errUnm)
	}
	fmt.Println(tmp)

}