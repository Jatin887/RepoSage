package main

import (
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/initializers"
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/models"
)



func main() {
	app := initializers.App()
	env := app.Env
	initializers.ConnectToDB(env)
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Repo{})
	initializers.DB.AutoMigrate(&models.UserRepo{})
}
