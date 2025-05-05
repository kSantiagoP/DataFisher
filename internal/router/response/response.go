package response

import "github.com/gin-gonic/gin"

func SendSuccess(c *gin.Context, data interface{}) {
	c.Header("Content-type", "application/json")
	c.JSON(200, data)
}

func SendError(c *gin.Context, code int, msg string) {
	c.Header("Content-type", "application/json")
	c.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}
