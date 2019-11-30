package middleware

import (
	"time"

	"github.com/teera123/gin"
)

//Logging monitor all service traffic
func Logging(c *gin.Context) {
	t := time.Now()
	// Set example variable
	c.Set("request_id", c.Request.Header.Get("request_id"))
	logger.Infof("start request URL %s from %s", c.Request.RequestURI, c.ClientIP())
	// before request
	c.Next()
	res := c.Writer.(gin.ResponseWriter)
	logger.Infof("response URL %s at %s status %d", c.Request.RequestURI, time.Since(t), res.Status())

}
