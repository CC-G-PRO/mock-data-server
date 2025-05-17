package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GoogleLogin(c *gin.Context) {
	type Req struct {
		IdToken string `json:"id_token"`
	}

	var req Req
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Invalid filter request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": "JWT_TOKEN",
		"user": gin.H{
			"id":    "user_id",
			"name":  "홍길동",
			"email": "hong@khu.ac.kr",
		},
	})
}
