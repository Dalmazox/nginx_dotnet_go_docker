package main

import (
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.GET("/api/v1/banks/:id", func(c *gin.Context) {
		pathId := c.Param("id")
		id, err := strconv.Atoi(pathId)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "id must be a number",
			})
			return
		}

		if id%2 != 0 {
			c.JSON(404, gin.H{
				"error": "bank not found",
			})
			return
		}

		c.JSON(200, &Bank{
			BankId: uuid.NewString(),
			Name:   "Bank " + pathId,
		})
	})

	r.Run(":" + os.Getenv("PORT"))
}

type Bank struct {
	BankId string `json:"bankId"`
	Name   string `json:"name"`
}
