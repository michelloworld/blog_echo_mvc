package flash

import (
	"time"

	"github.com/labstack/echo"
)

func SetFlash(name string, message string, c echo.Context) {
	cookie := new(echo.Cookie)
	cookie.SetName(name)
	cookie.SetValue(message)
	c.SetCookie(cookie)
}

func GetFlash(name string, c echo.Context) string {
	cookie, err := c.Cookie(name)

	if err == nil {
		message := cookie.Value()

		// delete cookie
		oldCookie := new(echo.Cookie)
		oldCookie.SetName(name)
		oldCookie.SetValue("")
		oldCookie.SetExpires(time.Now())
		c.SetCookie(oldCookie)

		return message
	}
	return ""
}
