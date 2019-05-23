package v1

import (
	"blog/pkg/logging"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {

	logging.MsgInfo(logrus.WithField("A", "a"), "This is a test")

	c.JSON(200, gin.H{
		"message": "test",
	})
}
