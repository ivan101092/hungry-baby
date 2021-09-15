package google

type Domain struct {
}

type Repository interface {
	GetTokenFromWeb(redirectURL string, scopes []string) (string, error)
	SaveTokenFromWeb(redirectURL string, scopes []string, authCode, destinationPath string) (string, error)
	SaveRefreshToken(redirectURL string, scopes []string, token interface{}, destinationPath string) error
	GetGoogleProfile(token string) (map[string]interface{}, error)
	GetTokenInfo(token string) (map[string]interface{}, error)
	VerifyTokenScope(token string, allowedScopes []string) error
}
