package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-app/httpd/handler"
	"go-gin-app/platform/newsfeed"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()
	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsFeedGet(feed))
	r.POST("/newsfeed", handler.NewsFeedPost(feed))
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
