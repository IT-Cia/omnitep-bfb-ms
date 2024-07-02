package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParamInt64(c *gin.Context, key string) (int64, error) {
	raw := c.Param("id")
	if raw == "" {
		return 0, nil
	}

	param, err := strconv.ParseInt(raw, 10, 64)

	return param, err
}

func QueryInt64(c *gin.Context, key string) (int64, error) {
	raw := c.Query("id")
	if raw == "" {
		return 0, nil
	}

	param, err := strconv.ParseInt(raw, 10, 64)

	return param, err
}
