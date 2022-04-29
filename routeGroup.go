package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Use(RequestInfo())
	user := r.Group("/user")
	user.POST("/register", RegisterHandle())

	if err := r.Run(); err != nil {
		log.Fatalln(err.Error())
		return
	}
}

func RegisterHandle() func(context *gin.Context) {
	return func(context *gin.Context) {
		var user User
		err := context.ShouldBindJSON(&user)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		context.JSON(http.StatusOK, user)
	}
}

// 打印请求的中间价
func RequestInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method
		fmt.Println("请求path:", path)
		fmt.Printf("method: %v \n", method)
		context.Next()

		fmt.Println("状态码：", context.Writer.Status())
	}

}
