package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/DanielGregorini/go-api-gin/service"
)

var (
	authService service.AuthService = service.NewAuthService(os.Getenv("SECRET_KEY"))
)

// Rotas públicas que não exigem token
var publicRoutes = []struct {
	Method string
	Path   string
}{
	{"POST", "/login"},
	{"POST", "/users"},
	//{"GET", "/videos"},
	{"GET", "/users"},
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
	
		//to public routes
		for _, route := range publicRoutes {
			if c.Request.Method == route.Method && c.Request.URL.Path == route.Path {
				c.Next()
				return
			}
		}

		// if dont have token
		tokenStr := c.GetHeader("token")
		if !authService.IsTokenValid(tokenStr) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou expirado"})
			c.Abort()
			return
		}

		// if have token
		c.Next()
	}
}
