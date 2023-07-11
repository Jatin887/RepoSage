package controllers

import (
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/config"
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/helpers"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func RootHandler(c *gin.Context) {
	c.Writer.WriteString(`<div>
		<a href="/login/github/">LOGIN</a>
		<span style="margin: 0 20px;"></span>
		<a href="/logout/">LOGOUT</a>
	</div>`)
}

func GithubLoginHandler(c *gin.Context) {
	// Get the environment variable
	githubClientID := helpers.GetGithubClientID()

	// Create the dynamic redirect URL for login
	redirectURL := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s",
		githubClientID,
		fmt.Sprintf("%s/login/github/callback", config.BaseURL),
	)
	c.Redirect(http.StatusMovedPermanently, redirectURL)
}



func GithubCallbackHandler(c *gin.Context) {
	code := c.Request.URL.Query().Get("code")

	// Log the code parameter
	log.WithFields(log.Fields{
		"code": code,
	}).Info("Received Github OAuth code")

	githubAccessToken, _ := getGithubAccessToken(code)

	// Log the Github access token
	log.WithFields(log.Fields{
		"access_token": githubAccessToken,
	}).Debug("Received Github access token")

	helpers.GetGithubData(githubAccessToken)

	c.JSON(200, gin.H{
		"message": "You are sucessfully logged in your secert access_token is : " + githubAccessToken,
	})
}


func getGithubAccessToken(code string) (string, error) {
	clientID := helpers.GetGithubClientID()
	clientSecret := helpers.GetGithubClientSecret()

	// Define the request body as JSON
	requestBodyMap := map[string]string{
		"client_id":     clientID,
		"client_secret": clientSecret,
		"code":          code,
	}
	requestJSON, err := json.Marshal(requestBodyMap)
	if err != nil {
		return "", err
	}

	// Create a new POST request to obtain the access token
	req, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(requestJSON),
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Send the request and retrieve the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body and parse the JSON
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var tokenResp struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}
	if err := json.Unmarshal(respBody, &tokenResp); err != nil {
		return "", err
	}
	// Return the access token
	return tokenResp.AccessToken, nil
}