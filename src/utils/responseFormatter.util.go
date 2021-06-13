package utils

import (
	"github.com/gin-gonic/gin"
)

//*Response Formatter return a bunch of templated responses in JSON format with status code condition
func ResponseFormatter(code int, message string, err error, result map[string]interface{}, c *gin.Context) {
	if code < 400 {
		c.AbortWithStatusJSON(code, gin.H{
			"code":    code,
			"success": true,
			"message": message,
			"error":   nil,
			"data":    result,
		})
		return
	}
	c.JSON(code, gin.H{
		"code":    code,
		"success": false,
		"message": message,
		"error":   err,
		"data":    result,
	})
}
