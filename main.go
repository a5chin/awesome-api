package main

import (
	db_model "api/model"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/questions", func(c *gin.Context) {
		questins := db_model.GetAll()
		c.JSON(200, questins)
	})
	r.Run()
}
