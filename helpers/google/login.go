package google

import (
	"encoding/json"
	"errors"
	"hungry-baby/helpers/str"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	OauthGoogleURLAPI          = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	OauthGoogleTokenInfoURLAPI = "https://www.googleapis.com/oauth2/v2/tokeninfo?access_token="
)

// GetGoogleProfile ...
func GetGoogleProfile(token string) (res map[string]interface{}, err error) {
	response, err := http.Get(OauthGoogleURLAPI + token)
	if err != nil {
		return res, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(responseBody, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

// GetTokenInfo ...
func GetTokenInfo(token string) (res map[string]interface{}, err error) {
	response, err := http.Get(OauthGoogleTokenInfoURLAPI + token)
	if err != nil {
		return res, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(responseBody, &res)
	if err != nil {
		return res, err
	}

	return res, err
}

// VerifyTokenScope ...
func VerifyTokenScope(token string, allowedScopes []string) error {
	userScope, err := GetTokenInfo(token)
	if err != nil {
		return err
	}
	if userScope["scope"] == nil {
		return errors.New("empty_scope")
	}

	scopes := strings.Split(userScope["scope"].(string), " ")
	var totalScopes int
	for _, a := range allowedScopes {
		if str.Contains(scopes, a) {
			totalScopes++
		}
	}
	if totalScopes != len(allowedScopes) {
		return errors.New("invalid_scope")
	}

	return nil
}
