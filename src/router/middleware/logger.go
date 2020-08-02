package middleware

import (
	"github.com/gin-gonic/gin"
	"os"
)

func Logger() gin.HandlerFunc {

	return gin.LoggerWithWriter(os.Stdout)

}
