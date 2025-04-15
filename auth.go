package unifi

import (
	"net/http"
)

func (c Client) Login() (bool, error) {

	req, err := http.NewRequest("POST", c.baseURL+"/api/login", nil)
	if err != nil {
		return false, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	if resp.StatusCode == 200 {
		return true, nil
	}
	return false, nil
}

func (c Client) Logout() (bool, error) {
	return false, nil
}
