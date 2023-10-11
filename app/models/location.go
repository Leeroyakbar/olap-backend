package models

import (
	"github.com/leeroyakbar/olap-backend/app/configs/clients"
)

type Dim_Location struct {
	Id_location int    `json:"id" gorm:"primaryKey"`
	Pulau       string `json:"pulau" form:"pulau"`
	Provinsi    string `json:"provinsi" form:"provinsi"`
	Kab_kota    string `json:"kota" form:"kota"`
	Kecamatan   string `json:"kecamatan" form:"kecamatan"`
	Desa        string `json:"desa" form:"desa"`
}

func (data Dim_Location) Add() error {
	return clients.DATABASE.Model(&data).Create(&data).Error
}

func (Dim_Location) TableName() string {
	return "dim_location"
}
