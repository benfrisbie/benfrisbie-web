package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	renderer := multitemplate.NewRenderer()
	renderer.AddFromFiles("index", "assets/templates/base.html", "assets/templates/index.html")
	renderer.AddFromFiles("404", "assets/templates/base.html", "assets/templates/404.html")
	router.HTMLRender = renderer
	router.Static("/static", "./assets/static")

	// Home page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{})
	})

	// 404 not found page
	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "404", gin.H{})
	})

	log.Fatal(router.Run()) // listen and serve on 0.0.0.0:8080
}
