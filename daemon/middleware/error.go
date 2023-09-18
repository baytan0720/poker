package middleware

import (
	gin "github.com/baytan0720/gin-grpc"
	log "github.com/sirupsen/logrus"

	"poker/pkg/errno"
)

func ErrorHandle(c *gin.Context) {
	c.Next()
	if c.Errors.Len() > 0 {
		err := c.Errors.Last()
		c.Resp.SetField("ErrorCode", errno.Errno(err))
		c.Errors.Clear()
		log.Error(err.Error())
	}
}
