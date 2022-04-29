package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type User struct {
	Username string `form:"username"`
	Tel      string `form:"tel" `
	Password string `form:"password"`
}

type Student struct {
	Name    string `form:"name"`
	Classes string `form:"classes"`
}

func main() {
	r := gin.Default()

	//r.GET("/hello", func(context *gin.Context) {
	//	fmt.Printf("context.FullPath(): %v \n", context.FullPath())
	//	context.Writer.Write([]byte("hello world\n"))
	//})

	// http:localhost:8080/hello?name=davie
	r.Handle("POST", "/hello", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Printf("path: %v \n", path)
		name := context.DefaultQuery("name", "jalivv")
		fmt.Printf("name: %v \n", name)

		var user User
		err := context.ShouldBind(&user)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		fmt.Printf("user.Username: %v \n", user.Username)
		fmt.Printf("user.Password: %v \n", user.Password)

		context.Writer.Write([]byte(user.Username + "====" + user.Password))

	})

	r.Handle("POST", "/login", func(context *gin.Context) {
		path := context.FullPath()
		fmt.Printf("path: %v \n", path)
		username := context.PostForm("username")
		password := context.PostForm("password")

		fmt.Printf("username:%s password:%s\n", username, password)

	})

	r.GET("/bind", func(context *gin.Context) {
		fmt.Printf("context.FullPath(): %v \n", context.FullPath())
		var stu Student
		err := context.ShouldBindQuery(&stu)

		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		context.Writer.Write([]byte(stu.Name + "====" + stu.Classes))

	})

	r.POST("/addStudent", func(context *gin.Context) {
		var user User
		if err := context.ShouldBindJSON(&user); err != nil {
			log.Fatalln(err.Error())
			return
		}

		context.Writer.Write([]byte(user.Username))
	})

	if err := r.Run(); err != nil {
		log.Fatalln(err.Error())

	}

}
