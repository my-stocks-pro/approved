package client

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strconv"
	"encoding/json"
)

func (a *ApprovedType) BaseRequest(page string) (*BaseRespType, error) {

	a.Base.NewURL = fmt.Sprintf(a.Config.BaseURL,
		a.Base.CurrYear, a.Base.CurrMonth, "%20", a.Base.CurrYear, a.Base.CurrMonth, a.Base.CurrDay, page, strconv.Itoa(a.Base.PerPage))

	res, errRequest := a.NewRequest()
	if errRequest != nil {
		return nil, errRequest
	}

	baseResp := BaseRespType{}
	errUnmarshal := json.Unmarshal(res, &baseResp)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &baseResp, nil
}

func (a *ApprovedType) FullRequest(id string) (*DataImageType, error) {
	a.Base.NewURL = a.Config.ApiURL + id

	res, errRequest := a.NewRequest()
	if errRequest != nil {
		return nil, errRequest
	}

	resp := DataImageType{}
	errUnmarshal := json.Unmarshal(res, &resp)
	if errUnmarshal != nil {
		return nil, errUnmarshal
	}

	return &resp, nil

}

func (a *ApprovedType) NewRequest() ([]byte, error) {

	fmt.Println(a.Base.NewURL)

	req, err := http.NewRequest("GET", a.Base.NewURL, nil)
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
