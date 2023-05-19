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

func (client Client) GetIsParent(jmbg string) (bool, error) {
	requestURL := client.address + fmt.Sprintf("/isParent/%s", jmbg)
	httpReq, err := http.NewRequest(http.MethodGet, requestURL, nil)

	if err != nil {
		return false, errors.New("req err")
	}

	res, err := http.DefaultClient.Do(httpReq)

	if err != nil {
		return false, errors.New("error getting info")
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	var isParent string

	json.Unmarshal(body, &isParent)
	if err != nil {
		return false, errors.New("error decoding body")
	}

	if isParent == "true" {
		return true, nil
	} else {
		return false, nil
	}
}
