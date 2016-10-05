package controllers

import (
	"blog_echo/app/models"
	"blog_echo/libs/flash"
	"blog_echo/libs/view"
	"strconv"

	"github.com/labstack/echo"
)

type BlogController struct{}

func (ctrl *BlogController) Index(c echo.Context) error {
	blogs, _ := new(models.Blog).All()
	flash_error := flash.GetFlash("flash_error", c)
	flash_success := flash.GetFlash("flash_success", c)
	return c.Render(200, "blogs/index", view.E{"blogs": blogs, "flash_error": flash_error, "flash_success": flash_success})
}

func (ctrl *BlogController) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	blog := new(models.Blog).FindById(id)
	return c.Render(200, "blogs/show", view.E{"blog": blog})
}

func (ctrl *BlogController) Create(c echo.Context) error {
	return c.Render(200, "blogs/create", nil)
}

func (ctrl *BlogController) Store(c echo.Context) error {
	blog := models.Blog{
		Title: c.FormValue("title"),
		Body:  c.FormValue("body"),
	}

	if !blog.Validate() {
		return c.Redirect(301, "/")
	}

	_, err := blog.Save()

	if err != nil {
		panic(err)
	}

	flash.SetFlash("flash_success", "Successfully Create", c)
	return c.Redirect(301, "/")
}

func (ctrl *BlogController) Edit(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	blog := new(models.Blog).FindById(id)
	return c.Render(200, "blogs/edit", view.E{"blog": blog})
}

func (ctrl *BlogController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	blog := models.Blog{
		Id:    id,
		Title: c.FormValue("title"),
		Body:  c.FormValue("body"),
	}

	_, err := blog.Update()

	if err != nil {
		panic(err)
	}

	flash.SetFlash("flash_success", "Successfully Update", c)
	return c.Redirect(301, "/")
}

func (ctrl *BlogController) Destroy(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	blog := new(models.Blog).FindById(id)

	_, err := blog.Delete()

	if err != nil {
		panic(err)
	}

	flash.SetFlash("flash_success", "Successfully Delete", c)
	return c.Redirect(301, "/")
}
