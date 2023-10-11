package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leeroyakbar/olap-backend/app/models"
)

func AddHotspot(ctx *fiber.Ctx) error {
	var data models.Fact_Hotspot

	if err := ctx.BodyParser(&data); err != nil {
		return ResponseJson(ctx, http.StatusBadRequest, err.Error())
	}

	if err := data.Add(); err != nil {
		return ResponseJson(ctx, http.StatusInternalServerError, err.Error())
	}

	return ResponseJson(ctx, http.StatusOK, "ok")
}

func AddLocation(ctx *fiber.Ctx) error {
	var data models.Dim_Location

	if err := ctx.BodyParser(&data); err != nil {
		return ResponseJson(ctx, http.StatusBadRequest, err.Error())
	}

	if err := data.Add(); err != nil {
		return ResponseJson(ctx, http.StatusInternalServerError, err.Error())
	}

	return ResponseJson(ctx, http.StatusOK, "ok")
}

func AddTime(ctx *fiber.Ctx) error {
	var data models.Dim_Time

	if err := ctx.BodyParser(&data); err != nil {
		return ResponseJson(ctx, http.StatusBadRequest, err.Error())
	}

	if err := data.Add(); err != nil {
		return ResponseJson(ctx, http.StatusInternalServerError, err.Error())
	}

	return ResponseJson(ctx, http.StatusOK, "ok")
}

func AddConfidence(ctx *fiber.Ctx) error {
	var data models.Dim_Confidence

	if err := ctx.BodyParser(&data); err != nil {
		return ResponseJson(ctx, http.StatusBadRequest, err.Error())
	}

	if err := data.Add(); err != nil {
		return ResponseJson(ctx, http.StatusInternalServerError, err.Error())
	}

	return ResponseJson(ctx, http.StatusOK, "ok")
}

func AddSatelite(ctx *fiber.Ctx) error {
	var data models.Dim_Satelite

	if err := ctx.BodyParser(&data); err != nil {
		return ResponseJson(ctx, http.StatusBadRequest, err.Error())
	}

	if err := data.Add(); err != nil {
		return ResponseJson(ctx, http.StatusInternalServerError, err.Error())
	}

	return ResponseJson(ctx, http.StatusOK, "ok")
}
