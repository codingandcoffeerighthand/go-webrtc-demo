package app

import (
	"fmt"
	"time"

	"github.com/codingandcoffeerighthand/go-webrtc-demo/internal/biz"
	"github.com/codingandcoffeerighthand/go-webrtc-demo/internal/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type app struct {
	host    string
	port    int
	authBiz handler.AuthBiz
}

func NewApp(host string, port int, secretKey string) *app {
	return &app{
		host:    host,
		port:    port,
		authBiz: biz.NewAuthBiz(secretKey),
	}
}
func (a *app) Start() {
	// Start the HTTP server
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/login", handler.LoginHandler(a.authBiz))

	r.Run(fmt.Sprintf("%s:%d", a.host, a.port))
	fmt.Printf("Server started at %s:%d\n", a.host, a.port)
}
