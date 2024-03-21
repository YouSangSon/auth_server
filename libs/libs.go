package libs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// this is what the data would look like, after the convert
// you can use this reference to get the data and save to db for example.
type GooglePayload struct {
	SUB           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
}

func ConvertToken(accessToken string) (*GooglePayload, error) {
	// call http request to this URL, this is a valid
	// URL provided from google to convert access token into
	// valid user data
	resp, httpErr := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/userinfo?access_token=%s", accessToken))
	if httpErr != nil {
		return nil, httpErr
	}

	// clean up when this function returns (destroyed)
	defer resp.Body.Close()

	// Reads the entire HTTP body from resp.Body using ioutil.ReadAll.
	// If any error occurs during the read operation, it is
	// returned as bodyErr. Otherwise it is stored in the respBody variable.
	respBody, bodyErr := ioutil.ReadAll(resp.Body)
	if bodyErr != nil {
		return nil, bodyErr
	}

	// Unmarshal raw response body to a map
	var body map[string]interface{}
	if err := json.Unmarshal(respBody, &body); err != nil {
		return nil, err
	}

	// If json body containing error,
	// then the token is indeed invalid. return invalid token err
	if body["error"] != nil {
		return nil, errors.New("Invalid token")
	}

	// Bind JSON into struct
	var data GooglePayload
	err := json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
