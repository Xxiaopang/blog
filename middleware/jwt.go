package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/config"
	"project/utils"
	"strings"
	"time"
)

func JWTAuto() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			utils.Response(c, http.StatusUnauthorized,
				nil,
				"请提供JWT Token!",
			)
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, "$", 2)
		if !(len(parts) == 2 && parts[0] == config.Header) {
			utils.Response(c, http.StatusUnauthorized, nil, "Invalid token type!")
			c.Abort()
			return
		}
		myClaims, ok := utils.CheckToken(parts[1])
		if !ok {
			utils.Response(c, http.StatusUnauthorized,
				nil,
				"Invalid token!",
			)
			c.Redirect(http.StatusMovedPermanently, "/api/login")

		}
		uid := myClaims.UID
		if time.Now().Unix() > myClaims.ExpiresAt.Unix() {
			utils.Response(c, http.StatusUnauthorized, nil, "Token timeout")
			c.Abort()
			return
		}
		newToken, err := utils.RefreshClaims(myClaims)
		if err != nil {
			utils.Response(c, http.StatusInternalServerError, nil, "Token refresh failed")
		} else {
			c.Header("Authorization", config.Header+"$"+newToken)
		}

		c.Set("uid", uid)
		c.Next()

	}
}
