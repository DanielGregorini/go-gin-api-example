package middleware

import (
	"net/http"
	"os"

	"github.com/DanielGregorini/go-api-gin/service"
	"github.com/gin-gonic/gin"
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
	{"GET", "/videos"},
	{"GET", "/users"},
}

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		for _, route := range publicRoutes {
			if context.Request.Method == route.Method && context.Request.URL.Path == route.Path {
				context.Next()
				return
			}
		}

		token := authService.ExtractTokenFromContext(context)

		if !authService.IsTokenValid(token) {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou expirado"})
			context.Abort()
			return
		}

		context.Next()
	}
}
