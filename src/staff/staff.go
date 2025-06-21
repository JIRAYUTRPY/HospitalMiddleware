package staff

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Staff struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Staff registration endpoint",
	})
}

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Staff login endpoint",
	})
}
