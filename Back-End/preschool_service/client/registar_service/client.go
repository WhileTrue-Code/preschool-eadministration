package registar_service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	address string
}

func NewClient(host, port string) Client {
	return Client{
		address: fmt.Sprintf("http://%s:%s", host, port),
	}
}

func (client Client) GetIsParent(jmbg, authToken string) (bool, error) {
	requestURL := client.address + fmt.Sprintf("/isParent/%s", jmbg)
	httpReq, err := http.NewRequest(http.MethodGet, requestURL, nil)
	httpReq.Header.Add("Authorization", authToken)

	if err != nil {
		return false, errors.New("req err")
	}

	res, err := http.DefaultClient.Do(httpReq)

	if err != nil {
		return false, errors.New("error getting info")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	var isParent bool

	json.Unmarshal(body, &isParent)
	if err != nil {
		return false, errors.New("error decoding body")
	}

	if isParent {
		return true, nil
	} else {
		return false, nil
	}
}

//func (client Client) GetChildren(jmbgRoditelja string) ([]data.User, error) {
//	requestURL := client.address + fmt.Sprintf("/children/%s", jmbgRoditelja)
//	httpReq, err := http.NewRequest(http.MethodGet, requestURL, nil)
//
//	var user []data.User
//
//	if err != nil {
//		return user, errors.New("req err")
//	}
//
//	res, err := http.DefaultClient.Do(httpReq)
//
//	if err != nil {
//		return user, errors.New("error getting info from registrar get children ")
//	}
//	defer res.Body.Close()
//
//	var deca []string
//
//	err = json.NewDecoder(res.Body).Decode(&deca)
//	if err != nil {
//		return user, errors.New("error decoding body get children")
//	}
//
//	return user, nil
//}
