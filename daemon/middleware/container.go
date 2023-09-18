package middleware

import (
	gin "github.com/baytan0720/gin-grpc"
	"poker/daemon/internal/manager"
)

func ValidName(c *gin.Context) {
	name := c.Req.GetField("Name").(string)
	err := manager.ValidateName(name)
	if err != nil {
		c.AbortWithError(err)
		return
	}
}
