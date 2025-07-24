package middleware

import (
    "log"
    "github.com/gin-gonic/gin"
)

func PrintLogger() gin.HandlerFunc {
	return func(context *gin.Context) {
		clientIP := context.ClientIP()

		log.Printf("-----------------------")
		
		log.Printf("Client IP: %s", clientIP)

		log.Printf("Request: %s %s", context.Request.Method, context.Request.URL.Path)
		context.Next()

		log.Printf("Response status: %d", context.Writer.Status())

		log.Printf("-----------------------")
	}
}