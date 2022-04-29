package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	user := r.Group("/user")
	user.POST("/register", func(context *gin.Context) {
		var user User
		err := context.ShouldBindJSON(&user)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}

		context.JSON(http.StatusOK, user)
	})

	if err := r.Run(); err != nil {
		log.Fatalln(err.Error())
		return
	}
}
