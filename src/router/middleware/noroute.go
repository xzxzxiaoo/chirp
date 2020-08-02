package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  1,
		"code":    404,
		"message": "Not Found",
	})
}
