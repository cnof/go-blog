package v1

import (
	"blog/models"
	"blog/pkg/logging"
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func Test(c *gin.Context) {

	validate := validator.New()

	var user models.User
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "",
		})
	}

	err = validate.Struct(user)

	if err != nil {
		return
	}

	logging.MsgInfo(logrus.WithField("A", "a"), "This is a test")

	c.JSON(http.StatusOK, gin.H{
		"message": "",
	})
}
