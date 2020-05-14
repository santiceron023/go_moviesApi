package middleare

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//func BenchMark() (func(c *gin.Context)) {
func BenchMark() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		elapsed := time.Since(start)
		fmt.Println("response took %v",elapsed)
	}
}