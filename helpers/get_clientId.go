package helpers

import (
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/initializers"
)


func GetGithubClientID() string {
	app := initializers.App()
	env := app.Env

	githubClientID := env.ClientID

	return githubClientID
}

func GetGithubClientSecret() string {
	app := initializers.App()
	env := app.Env

	githubClientSecret := env.ClientSecret

	return githubClientSecret
}