package main

import (
	"net/http"

	"github.com/orestonce/ymd/ymdGin"
)

var db = make(map[string]string)

func setupRouter() *ymdGin.Engine {
	// Disable Console Color
	// ymdGin.DisableConsoleColor()
	r := ymdGin.Default()

	// Ping test
	r.GET("/ping", func(c *ymdGin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *ymdGin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, ymdGin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, ymdGin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses ymdGin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(ymdGin.BasicAuth(ymdGin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", ymdGin.BasicAuth(ymdGin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *ymdGin.Context) {
		user := c.MustGet(ymdGin.AuthUserKey)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, ymdGin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
