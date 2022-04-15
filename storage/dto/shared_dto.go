package dto

import (
	"github.com/gin-gonic/gin"
	"sfeir/global"
)

type PaginationDto struct {
	Limit int64 `json:"limit" bson:"limit"`
	Skip  int64 `json:"skip" bson:"skip"`
}

func NewPaginationDto(c *gin.Context) PaginationDto {
	page := global.GetInt64ValueFromQueryOrDefault(c, global.Page, global.DefaultPage)
	limit := global.GetInt64ValueFromQueryOrDefault(c, global.Limit, global.DefaultLimit)

	return PaginationDto{
		Skip:  page * limit,
		Limit: limit,
	}
}
