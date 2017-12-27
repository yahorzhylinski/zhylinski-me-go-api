package main

import (
	"github.com/labstack/echo"
	"net/http"
	. "./db"
	"./models"
)

func getProfile(context echo.Context) error {
	userProfile, err := DB().FindByPrimaryKeyFrom(models.ProfileTable, context.Param("id"))
	if err != nil {
		context.Logger().Error("FOO")
	}

	return context.String(http.StatusOK, userProfile.String())
}

func main() {
	e := echo.New()
	e.GET("/api/profiles/:id", getProfile)
	e.Logger.Fatal(e.Start(":9090"))
}
