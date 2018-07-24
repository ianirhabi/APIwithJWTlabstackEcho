package router

import (
	"belajar/echolabstack/JWTToken/auth"
	"belajar/echolabstack/JWTToken/src/absen"
	"belajar/echolabstack/JWTToken/src/usr"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Router() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("/absen", absen.Absen)

	// untuk absen
	e.GET("/", usr.Getuser)

	v := e.Group("/auth")
	v.POST("/login", auth.Login)

	e.Logger.Fatal(e.Start(":1323"))
}
