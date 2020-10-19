package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/kainhuck/gold/pkg/config"
)

// GetPage get page parameters
func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * config.Collection.App.PageSize
	}

	return result
}
