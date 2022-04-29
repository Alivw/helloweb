package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

type Resp struct {
	Code    int
	Message string
	Data    interface{}
}

func main1() {
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

	r.POST("/helloJson", func(context *gin.Context) {
		var user User
		err := context.ShouldBindJSON(&user)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"code":    1,
			"message": "ok",
			"data":    user,
		})

	})

	r.POST("/helloStruct", func(context *gin.Context) {

		var user User
		err := context.ShouldBindJSON(&user)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		resp := Resp{Code: 1, Message: "ok", Data: user}
		context.JSON(http.StatusOK, &resp)

	})

	// 设置html 目录
	r.LoadHTMLGlob("./html/*")

	// 设计静态资源目录
	r.Static("/img", "./img")
	r.GET("/hello/html", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"user":  "jalivv",
			"Title": "gin教程",
		})
	})
	if err := r.Run(); err != nil {
		log.Fatalln(err.Error())

	}

}
