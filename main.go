package main

import (
    "embed"
    "github.com/gin-gonic/gin"
    "io/fs"
    "net/http"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()
    r.LoadHTMLGlob("templates/*")
    staticFS, _ := fs.Sub(staticFiles, "static")
    r.StaticFS("/static", http.FS(staticFS))
    r.GET("/ping", func(c *gin.Context) { c.HTML(http.StatusOK, "ping.html", nil) })
    r.NoRoute(func(c *gin.Context) { c.HTML(http.StatusNotFound, "404.html", nil) })
    r.RunTLS(":443", "fullchain.pem", "privkey.pem")
}