// +build go1.8

package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"github.com/orestonce/ymd/ymdGin"
)

func main() {
	router := ymdGin.Default()
	router.GET("/", func(c *ymdGin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			log.Fatal("Server Close:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpect")
		}
	}

	log.Println("Server exiting")
}
