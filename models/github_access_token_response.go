package models
import(
	"gorm.io/gorm"
)
type GithubAccessTokenResponse struct {
	gorm.Model
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}