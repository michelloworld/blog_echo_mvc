package main

import (
	"blog_echo/app/configs"
	"blog_echo/app/controllers"
	"blog_echo/libs/template"
	"html/template"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Static("/assets", "./public")
	e.Use(middleware.Logger())

	// db
	configs.InitDB("mysql", "root:@/gin_blog?parseTime=True")

	// template
	t := eztemplate.New()
	t.TemplateFuncMap = templateFunc()
	e.SetRenderer(&t)

	// routes
	e.GET("/", (&controllers.BlogController{}).Index)
	e.GET("/blog/new", (&controllers.BlogController{}).Create)

	e.GET("/blog/:id", (&controllers.BlogController{}).Show)
	e.POST("/blog", (&controllers.BlogController{}).Store)

	e.GET("/blog/:id/edit", (&controllers.BlogController{}).Edit)
	e.POST("/blog/:id/update", (&controllers.BlogController{}).Update)

	e.GET("/blog/:id/delete", (&controllers.BlogController{}).Destroy)

	// server
	e.Run(fasthttp.New(":9000"))
}

func templateFunc() template.FuncMap {
	return template.FuncMap{
		"nl2br": func(text string) template.HTML {
			return template.HTML(strings.Replace(template.HTMLEscapeString(text), "\n", "<br>", -1))
		},
		"dateFormat": func(t time.Time) string {
			return t.Format("2 Jan 2006 15:04:05")
		},
	}
}
