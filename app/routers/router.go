package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leeroyakbar/olap-backend/app/handlers"
)

func Router(app *fiber.App) {
	base := app.Group("/api")

	base.Get("/query/", handlers.GetHeader)
	base.Get("/query/:dimension", handlers.Query)

	// lokasi
	location := base.Group("/location")
	location.Post("", handlers.AddLocation)
	// location.Get("", handlers.QueryLocation)

	// waktu
	time := base.Group("/time")
	time.Post("", handlers.AddTime)

	// satelit
	satelite := base.Group("/satelite")
	satelite.Post("", handlers.AddSatelite)

	// confidence
	confidence := base.Group("/confidence")
	confidence.Post("", handlers.AddConfidence)

	// lokasi
	hotspot := base.Group("/hotspot")
	hotspot.Post("", handlers.AddHotspot)

}
