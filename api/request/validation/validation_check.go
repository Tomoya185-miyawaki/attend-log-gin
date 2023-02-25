package validation

import (
	"github.com/gin-gonic/gin"
)

func ValidationCheck(
	c *gin.Context,
	request interface{},
) (bool, error) {
	if err := c.BindJSON(&request); err != nil {
		return false, err
	}
	return true, nil
}
