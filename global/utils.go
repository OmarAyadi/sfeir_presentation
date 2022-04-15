package global

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func ToStr(value any) string {
	return fmt.Sprintf("%v", value)
}

func InvalidIDErr(value any) error {
	invalidIDTemplate := fmt.Sprintf("invalid id '%v'", value)
	return errors.New(invalidIDTemplate)
}

func GetInt64ValueFromQueryOrDefault(c *gin.Context, key string, defaultValue int64) int64 {
	value := c.Query(key)
	convertedValue, err := strconv.ParseInt(value, 10, 64)
	if value == "" || err != nil {
		return defaultValue
	}
	return convertedValue
}
