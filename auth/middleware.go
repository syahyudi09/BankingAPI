	package auth

	import (
		"net/http"
		"strings"

		"github.com/dgrijalva/jwt-go"
		"github.com/gin-contrib/sessions"
		"github.com/gin-gonic/gin"
	)

	// Middleware adalah antarmuka (interface) untuk middleware autentikasi.
	type Middleware interface {
		AuthMiddleware() gin.HandlerFunc
		AuthCustomerMiddleware() gin.HandlerFunc
	}

	type middleware struct {
		authService Service
	}

	// AuthMiddleware adalah fungsi untuk memeriksa token autentikasi dari header "Authorization".
	// Fungsi ini harus diimplementasikan oleh tipe yang memenuhi antarmuka Middleware.
	func (m *middleware) AuthMiddleware() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			authHeader := ctx.GetHeader("Authorization")

			if !strings.Contains(authHeader, "Bearer") {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})
				ctx.Abort()
				return
			}

			stringToken := ""
			tokenString := strings.Split(authHeader, " ")
			if len(tokenString) == 2 {
				stringToken = tokenString[1]
			}

			token, err := m.authService.ValidateToken(stringToken)
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})
				ctx.Abort()
				return
			}

			_, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthorized",
				})
				ctx.Abort()
				return
			}

			ctx.Next()
		}
	}

	func (m *middleware) AuthCustomerMiddleware() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			session := sessions.Default(ctx)
			customerIDsession := session.Get("customerID")
	
			if customerIDsession == nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"error": "Unauthorized",
				})
				ctx.Abort()
				return
			}
			ctx.Next()
		}
	}
	
	// NewMiddleware adalah fungsi pembuat untuk membuat instance Middleware.
func NewMiddleware(authService Service) Middleware {
	return &middleware{
			authService: authService,
		}
	}
