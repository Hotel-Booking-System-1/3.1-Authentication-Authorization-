package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/mubarik-siraji/booking-system/infra"
)

// AuthMiddleware validates the JWT token provided in the Authorization header
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Authorization header is missing",
				"data":    nil,
			})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Authorization header format must be Bearer {token}",
				"data":    nil,
			})
			c.Abort()
			return
		}

		tokenString := parts[1]
		config := infra.Configurations

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(config.AccsessJwtToKenSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Invalid or expired token",
				"data":    nil,
			})
			c.Abort()
			return
		}

		// Extract claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Ensure token is not a refresh token
			if isRefresh, ok := claims["isRefreshToken"].(bool); ok && isRefresh {
				c.JSON(http.StatusUnauthorized, gin.H{
					"status":  "error",
					"message": "Refresh token cannot be used to access this resource",
					"data":    nil,
				})
				c.Abort()
				return
			}

			// Store user information in the context
			c.Set("userEmail", claims["sub"])
			c.Set("userRole", claims["role"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Invalid token claims",
				"data":    nil,
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// AuthorizeRoles ensures the user making the request has one of the allowed roles
func AuthorizeRoles(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Unauthorized: role information not found",
			})
			c.Abort()
			return
		}

		roleStr := fmt.Sprintf("%v", userRole)
		roleStrLower := strings.ToLower(roleStr)

		isAllowed := false
		for _, role := range allowedRoles {
			if strings.ToLower(role) == roleStrLower {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  "error",
				"message": "Forbidden: you don't have permission to access this resource",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
