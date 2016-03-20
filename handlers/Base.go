package handlers

import (
	"net/http"
	"fmt"
	
	"github.com/gin-gonic/gin"
)

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}

	c.JSON(code, resp)
	c.AbortWithStatus(code)
}


func jwtVerifier() gin.HandlerFunc {
	return func(c *gin.Context) {

		appToken := c.Request.Header.Get("Authorization")

		if appToken == "" {
			respondWithError(http.StatusForbidden, "Authorization header is required", c)
		} else {
			respondWithError(http.StatusBadRequest, fmt.Sprintf("Invalid token: %s", appToken), c)
		}
	}
}