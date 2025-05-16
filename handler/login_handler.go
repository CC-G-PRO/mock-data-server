package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GoogleLogin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"access_token": "JWT_TOKEN",
		"user": gin.H{
			"id":    "user_id",
			"name":  "홍길동",
			"email": "hong@khu.ac.kr",
		},
	})
}
