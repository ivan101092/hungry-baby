package google

import (
	"encoding/json"
	"errors"
	googleBusiness "hungry-baby/businesses/google"
	"hungry-baby/helpers/interfacepkg"
	"hungry-baby/helpers/str"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/rs/xid"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/people/v1"
)

// OauthCredential ...
type OauthCredential struct {
	Key         string
	RedirectURL string
}

func NewGoogle(key, redirectURL string) googleBusiness.Repository {
	return &OauthCredential{
		Key:         key,
		RedirectURL: redirectURL,
	}
}

var (
	// CalendarScopes ...
	CalendarScopes = []string{
		people.UserinfoProfileScope, people.UserinfoEmailScope, calendar.CalendarScope,
	}
)

const (
	OauthGoogleURLAPI          = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	OauthGoogleTokenInfoURLAPI = "https://www.googleapis.com/oauth2/v2/tokeninfo?access_token="
)

// GetScope ...
func GetScope(types string) []string {
	return CalendarScopes
}

// GetClient Retrieve a token, saves the token, then returns the generated client.
func GetClient(config *oauth2.Config, tokenString string) *http.Client {
	tok, err := tokenFromString(tokenString)
	if err != nil {
		return nil
	}
	return config.Client(context.Background(), tok)
}

// tokenFromString ...
func tokenFromString(tokenString string) (*oauth2.Token, error) {
	token := strings.NewReader(tokenString)
	tok := &oauth2.Token{}
	err := json.NewDecoder(token).Decode(tok)
	return tok, err
}

// GetTokenFromWeb Request a token from the web, then returns the retrieved token.
func (cred *OauthCredential) GetTokenFromWeb(redirectURL string, scopes []string) (string, error) {
	key := strings.Replace(cred.Key, "{redirect_url}", cred.RedirectURL, 1)
	b := []byte(key)
	config, err := google.ConfigFromJSON(b, scopes...)
	if err != nil {
		return "", err
	}

	return config.AuthCodeURL("state-token", oauth2.AccessTypeOffline) + "&prompt=select_account", nil
}

// SaveTokenFromWeb Request a token from the web, then returns the retrieved token.
func (cred *OauthCredential) SaveTokenFromWeb(redirectURL string, scopes []string, authCode, destinationPath string) (string, error) {
	res := destinationPath + xid.New().String() + ".json"

	key := strings.Replace(cred.Key, "{redirect_url}", cred.RedirectURL, 1)
	b := []byte(key)
	config, err := google.ConfigFromJSON(b, scopes...)
	if err != nil {
		return "", err
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return "", err
	}

	f, err := os.OpenFile(res, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return "", err
	}
	defer f.Close()
	json.NewEncoder(f).Encode(tok)

	return res, nil
}

// SaveRefreshToken Request a token from the web, then returns the retrieved token.
func (cred *OauthCredential) SaveRefreshToken(redirectURL string, scopes []string, token interface{}, destinationPath string) error {
	key := strings.Replace(cred.Key, "{redirect_url}", cred.RedirectURL, 1)
	b := []byte(key)
	config, err := google.ConfigFromJSON(b, scopes...)
	if err != nil {
		return err
	}

	var oauthToken *oauth2.Token
	interfacepkg.Convert(token, &oauthToken)

	tokenSource := config.TokenSource(oauth2.NoContext, oauthToken)
	tok, err := tokenSource.Token()
	if err != nil {
		return err
	}

	f, err := os.OpenFile(destinationPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	json.NewEncoder(f).Encode(tok)

	return nil
}

// GetGoogleProfile ...
func (cred *OauthCredential) GetGoogleProfile(token string) (res map[string]interface{}, err error) {
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
func (cred *OauthCredential) GetTokenInfo(token string) (res map[string]interface{}, err error) {
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
func (cred *OauthCredential) VerifyTokenScope(token string, allowedScopes []string) error {
	userScope, err := cred.GetTokenInfo(token)
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
