package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/opensourceway/app-robot-server/global"
	"github.com/opensourceway/app-robot-server/models"
	"github.com/opensourceway/app-robot-server/models/response"
)

func illegalToken(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusUnauthorized, response.BaseResponse{Code: code, Msg: msg})
	c.Abort()
}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get the token from the request header
		token := c.Request.Header.Get(global.TokenKey)
		if token == "" {
			illegalToken(c, global.UnauthorizedCode, global.EmptyTokenErrMsg)
			return
		}
		j := models.NewJwt()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == models.TokenExpired {
				illegalToken(c, global.UnauthorizedCode, global.TokenHasExpiredMsg)
				return
			}
			illegalToken(c, global.UnauthorizedCode, global.EmptyTokenErrMsg)
			return
		}
		//TODO: judge the server cache token is validate
		// ...
		c.Set(global.CKGToken, claims)
		c.Next()
	}
}
