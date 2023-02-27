package validation

import (
	"github.com/gin-gonic/gin"
)

func ValidationJsonCheck(
	c *gin.Context,
	request interface{},
) error {
	if err := c.BindJSON(&request); err != nil {
		return err
	}
	return nil
}
