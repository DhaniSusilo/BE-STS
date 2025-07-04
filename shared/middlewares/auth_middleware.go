package middlewares

import (
	"context"
	"learning-backend/domains/models/dto"
	"learning-backend/shared/constant"
	"learning-backend/shared/models/responses"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// ✅ 1. Define a custom context key type to avoid staticcheck warning
type contextKey string

// ✅ 2. Define the key constant for storing user info in context
const AuthUserKey contextKey = "auth-user"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken := strings.ReplaceAll(c.Request.Header.Get("Authorization"), "Bearer ", "")

		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			return constant.JWT_SECRET, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.BasicResponse{Error: "Unauthorized"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.BasicResponse{Error: "Unauthorized"})
			return
		}

		authUser := &dto.AuthAccessToken{}
		if tokenID, ok := claims["token_id"].(string); ok {
			authUser.Id = tokenID
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.BasicResponse{Error: "Invalid token ID"})
			return
		}

		if IsTokenRevoked(authUser.Id) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responses.BasicResponse{
			Error: "Token has been revoked",
		})
		return
}
		// Add more claims parsing as needed, like user_id or role

		// ✅ 5. Use the typed context key instead of raw string
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), AuthUserKey, authUser))

		c.Next()
	}
}

var revokedTokens sync.Map

// Add token to revoked list
func RevokeToken(tokenID string, expiresAt time.Time) {
	revokedTokens.Store(tokenID, expiresAt)
}

// Check if token is revoked
func IsTokenRevoked(tokenID string) bool {
	val, ok := revokedTokens.Load(tokenID)
	if !ok {
		return false
	}

	expTime := val.(time.Time)
	if expTime.Before(time.Now()) {
		// Auto-cleanup if expired
		revokedTokens.Delete(tokenID)
		return false
	}

	return true
}
