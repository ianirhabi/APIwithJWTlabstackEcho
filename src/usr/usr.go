package usr

import (
	"net/http"

	"github.com/labstack/echo"
)

func Getuser(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}
