package main

import (
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/api/controllers"
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/api/middelware"
	"balkanid-summer-internship-vit-vellore-2023-Jatin887/initializers"
	"fmt"
	"github.com/gin-gonic/gin"
)


func main(){
	r := gin.Default()
	app := initializers.App()
	env := app.Env
	initializers.ConnectToDB(env)
	initializers.InitLogger()
	r.GET("/",controllers.RootHandler)
	r.GET("/login/github/", controllers.GithubLoginHandler)
	r.GET("/login/github/callback",controllers.GithubCallbackHandler)
	r.GET("/dashboard",middelware.GithubAuthMiddleware(), controllers.DashboardHandler)
	r.GET("/download",middelware.GithubAuthMiddleware(),controllers.CsvHandler)
	r.GET("/logout",controllers.HandleSignout)
	fmt.Println("[ UP ON PORT 8080 ]")
	r.Run() 
}