package main

import (
	"errors"
	"github.com/labstack/echo/v4"
	"goproxy2/assets"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

type AdminTemplate struct {
	templates *template.Template
}

func (t *AdminTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func parseGlob(prefix string) (*template.Template, error) {
	var t *template.Template
	for _, name := range assets.AssetNames() {
		if strings.HasPrefix(name, prefix) {
			content, err := assets.Asset(name)
			if err != nil {
				return nil, err
			}
			var tmpl *template.Template
			if t == nil {
				t = template.New(name)
			}
			if name == t.Name() {
				tmpl = t
			} else {
				tmpl = t.New(name)
			}
			_, err = tmpl.Parse(string(content))
			if err != nil {
				return nil, err
			}
		}
	}
	return t, nil
}

func startAdmin() {
	Init()
	e := echo.New()
	e.HTTPErrorHandler = func(err error, ctx echo.Context) {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code": 500,
				"msg":  "请求失败" + err.Error(),
			})
		}
	}
	e.Renderer = &AdminTemplate{
		templates: template.Must(parseGlob("resources/views/")),
	}
	e.POST("/admin/add-ip", func(ctx echo.Context) error {
		remoteIp := ctx.FormValue("from")
		proxyIp := ctx.FormValue("to")
		if remoteIp == "" {
			remoteIp = GetIpFromRemoteAddr(ctx.Request().RemoteAddr)
		}
		if proxyIp == "" {
			return errors.New("proxyIp must not be empty")
		}
		AddIp(remoteIp, proxyIp)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
		})
	})
	e.POST("/admin/delete-ip", func(ctx echo.Context) error {
		remoteIp := ctx.FormValue("from")
		if remoteIp == "" {
			remoteIp = GetIpFromRemoteAddr(ctx.Request().RemoteAddr)
		}
		DeleteIp(remoteIp)
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code": 0,
		})
	})
	e.GET("/", func(ctx echo.Context) error {
		return ctx.Render(http.StatusOK, "admin/index", map[string]interface{}{
			"items": ipDb,
		})
	})
	err := e.Start(":8080")
	if err != nil {
		log.Println("listen :8080 failed", err)
	}
}
