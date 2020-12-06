package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
		})
		e := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if e != nil {
			return
		}

		c.Next()
	}
}
