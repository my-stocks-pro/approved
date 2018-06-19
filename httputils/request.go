package httputils

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/my-stocks-pro/approved/new"
)

func (a *new.ApprovedType) NewRequest(url string, needCookie bool, needHeader bool) []byte {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	if needHeader == true {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.Token))
	}

	if needCookie == true {
		cookie := http.Cookie{Name: "session", Value: "s%3AFLsDQ0KkRmbbHJSFijJz_5VxQPCQI7Ol.t5LQWhFeOPA9qV2S0fqa6JBsFB0Rq%2BrxMDPc1URXyHE"}
		req.AddCookie(&cookie)
	}

	client := http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		fmt.Println(errResp)
	}

	body, errReadBody := ioutil.ReadAll(resp.Body)
	if errReadBody != nil {
		fmt.Println(errReadBody)
	}

	defer resp.Body.Close()

	fmt.Println(string(body))

	return body

	//tmp := ImageType{}
	//
	//errUnm := json.Unmarshal(body, &tmp)
	//if errUnm != nil {
	//	fmt.Println(errUnm)
	//}
	//fmt.Println(tmp)

}
