package middleware

import "github.com/teera123/gin"

//AcceptLang default is EN
func AcceptLang(c *gin.Context) {
	l := c.Request.Header.Get("Accept-Language")
	if l == "" {
		l = "en"
	}
	c.Request.Header.Set("accept_language", l)
	c.Next()
}
