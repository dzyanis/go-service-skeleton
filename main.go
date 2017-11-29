package main

import (
    "os"
    "log"
    "github.com/gin-gonic/gin"
    "net/http"
)

var (
    RELEASE = "unset"
    REPO = "unset"
    COMMIT = "unset"
)

func main() {
    logger := log.New(os.Stdout, "", log.LstdFlags)

    port := os.Getenv("SERVICE_PORT")
    if len(port) == 0 {
        logger.Fatal("Required parameter service port is not set")
    }

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })
    r.GET("/info", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "RELEASE": RELEASE,
            "REPO": REPO,
            "COMMIT": COMMIT,
        })
    })
    r.GET("/healthz", func(c *gin.Context) {
        c.String(http.StatusOK, http.StatusText(http.StatusOK))
    })

    err := r.Run("0.0.0.0:" + port)
    if err != nil {
        logger.Printf("ERR: %s", err.Error())
    }
}