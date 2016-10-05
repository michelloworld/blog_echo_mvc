package eztemplate

import (
	"html/template"
	"io"
	"path/filepath"

	"github.com/labstack/echo"
)

var templates map[string]*template.Template

type Template struct {
	// TemplatesDir holds the location of the templates
	TemplateDir string
	// Layout is the file name of the layout file
	Layout string
	// Ext is the file extension of the templates
	Ext string
	// FuncMap
	TemplateFuncMap map[string]interface{}
}

func New() Template {
	return Template{
		TemplateDir:     "app/views/",
		Layout:          "layouts/base",
		Ext:             ".html",
		TemplateFuncMap: nil,
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	templates = make(map[string]*template.Template)
	layout := t.TemplateDir + t.Layout + t.Ext
	view := t.TemplateDir + name + t.Ext
	templates[name] = template.Must(template.New(filepath.Base(layout)).Funcs(t.TemplateFuncMap).ParseFiles(layout, view))
	return templates[name].ExecuteTemplate(w, filepath.Base(t.Layout+t.Ext), data)
}
