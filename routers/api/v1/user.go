package v1

import (
	"blog/pkg/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func Test(c *gin.Context) {
	logger := logging.InitLogger("./runtime/logs/user.log", "info")

	logger.Info("this is a test",
		zap.String("test", "test"),
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "aaa",
	})
}
