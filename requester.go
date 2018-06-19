package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/dyninc/qstring"
)

type Query struct {
	ID   []string
	View string
}

func Req(newIDs []string) {


	query := &Query{
		ID: newIDs,
		View: "full",
	}
	q, errQ := qstring.MarshalString(query)
	if errQ != nil {
		fmt.Println(errQ)
	}

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

	//tmp := TestType{}

	//errUnm := json.Unmarshal(body, &tmp)
	//if errUnm != nil {
	//	fmt.Println(errUnm)
	//}

}