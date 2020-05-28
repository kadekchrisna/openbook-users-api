package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is for health check
func Ping(c *gin.Context) {

	// it's the same thing as line 16
	// c.JSON(http.StatusOK, map[string]interface{}{
	// 	"message": "OK",
	// })
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "OK",
	// })

	// Sending only string and status code as a response
	c.String(http.StatusOK, "PONG")
}
