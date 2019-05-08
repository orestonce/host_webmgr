package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/orestonce/ymd/ymdGin"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d%02d/%02d", year, month, day)
}

func main() {
	router := ymdGin.Default()
	router.Delims("{[{", "}]}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	router.LoadHTMLFiles("../../testdata/template/raw.tmpl")

	router.GET("/raw", func(c *ymdGin.Context) {
		c.HTML(http.StatusOK, "raw.tmpl", map[string]interface{}{
			"now": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC),
		})
	})

	router.Run(":8080")
}
