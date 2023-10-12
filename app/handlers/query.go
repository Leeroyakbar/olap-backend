package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leeroyakbar/olap-backend/app/helpers"
	"github.com/leeroyakbar/olap-backend/app/models/payload"
	"github.com/xdbsoft/olap"
)


func GetHeader(ctx *fiber.Ctx) error {
	var cube olap.Cube
	var err error
	var other payload.Other
	limit := ctx.QueryInt("limit", 100)
	offset := ctx.QueryInt("offset", 10)

	cube, err = helpers.SQLtoCube(other, limit, offset)
	if err != nil {
		return ResponseJson(ctx, http.StatusBadGateway, err.Error())
	}

	return ResponseJson(ctx, http.StatusOK, cube.Headers())
}

func Query(ctx *fiber.Ctx) error {
	var cube olap.Cube
	var err error

	dimension := ctx.Params("dimension")
	point := ctx.Query("point")
	dim := ctx.Query("dimension")
	limit := ctx.QueryInt("limit", 100)
	offset := ctx.QueryInt("offset", 10)

	location, time, other := helpers.GetPayload(ctx)

	if dimension == "location" || dimension == "time" {
		if location.Pulau != "" || time.Tahun != "" {
			if location.Pulau != "" && time.Tahun == "" {
				cube, err = helpers.CubeLocation(location, other)
				if err != nil {
					return ResponseJson(ctx, http.StatusBadGateway, err.Error())
				}
			}
			if location.Pulau == "" && time.Tahun != "" {
				cube, err = helpers.CubeTime(time, other)
				if err != nil {
					return ResponseJson(ctx, http.StatusBadGateway, err.Error())
				}
			}

			if location.Pulau != "" && time.Tahun != "" {
				cube, err = helpers.CubeTimeLocation(time, location, other)
				if err != nil {
					return ResponseJson(ctx, http.StatusBadGateway, err.Error())
				}
			}
		} else {
			cube, err = helpers.SQLtoCube(other, limit, offset)
			if err != nil {
				return ResponseJson(ctx, http.StatusBadGateway, err.Error())
			}
		}
	} else {
		cube, err = helpers.SQLtoCube(other, limit, offset)
		if err != nil {
			return ResponseJson(ctx, http.StatusBadGateway, err.Error())
		}
	}

	if point != "" && dim != "" {
		cube = cube.Slice(dimension, point)
		cube = cube.RollUp([]string{dim}, cube.Fields, helpers.Sum, []interface{}{0})
	} else {
		cube = cube.RollUp([]string{dimension}, cube.Fields, helpers.Sum, []interface{}{0})
	}

	return ResponseJson(ctx, http.StatusOK, cube.Rows())
}
