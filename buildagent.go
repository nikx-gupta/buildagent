package buildagent

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() error {

	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		var json GitEvent
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Print(json)
		c.Done()
	})

	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Print(json)

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})

	return r.Run()
}

type GitRepository struct {
	Url string `json:"url"`
}

type GitEvent struct {
	Repository GitRepository `json:"repository"`
}

type Login struct {
	User       string        `form:"user" json:"user" xml:"user"  binding:"required"`
	Password   string        `form:"password" json:"password" xml:"password" binding:"required"`
	Repository GitRepository `json:"repository"`
}
