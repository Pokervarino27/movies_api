package internal

import (
	"github.com/gin-gonic/gin"
	utils "github.com/pokervarino27/movies_api/utils"
)

//RedisConnection add to the context of gin a redis client interface
func RedisConnection(redis utils.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("redisConn", redis)
		c.Next()
	}
}
