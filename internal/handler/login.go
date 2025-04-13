package handler

import (
	"github.com/codingandcoffeerighthand/go-webrtc-demo/internal/biz"
	"github.com/gin-gonic/gin"
)

type AuthBiz interface {
	Login(biz.LoginData) (string, error)
}

func LoginHandler(authBiz AuthBiz) gin.HandlerFunc {
	return func(g *gin.Context) {
		var loginData biz.LoginData
		if err := g.ShouldBindJSON(&loginData); err != nil {
			g.JSON(400, gin.H{"error": err.Error()})
			return
		}
		token, err := authBiz.Login(loginData)
		if err != nil {
			g.JSON(500, gin.H{"error": err.Error()})
			return
		}
		g.JSON(200, gin.H{"token": token})
	}
}
