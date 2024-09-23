package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/houg/go-oauth2-resource/pkg/common"
	"github.com/houg/go-oauth2-resource/resource"
	"net/http"
)

func Oauth2ResourceMiddleware(scopes []string, grantTypes []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(scopes) == 0 {
			c.IndentedJSON(http.StatusUnauthorized, common.Result{
				Code: 401,
				Msg:  common.UNAUTHORIZED_ACCESS,
			})
		}
		accessToken, err := resource.Instance.ValidationBearerToken(c.Request)
		if err != nil {
			c.IndentedJSON(http.StatusUnauthorized, common.Result{
				Code: 401,
				Msg:  err.Error(),
			})
		}
		if !accessToken.HasScopes(scopes...) {
			c.IndentedJSON(http.StatusUnauthorized, common.Result{
				Code: 401,
				Msg:  common.UNAUTHORIZED_ACCESS,
			})
		}
		if len(grantTypes) > 0 && !accessToken.HasGrantType(grantTypes...) {
			c.IndentedJSON(http.StatusUnauthorized, common.Result{
				Code: 401,
				Msg:  common.UNAUTHORIZED_ACCESS,
			})
		}
		c.Set("accessToken", accessToken) // 从请求头中获取Token
		c.Next()
	}
}
