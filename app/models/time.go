package models

import (
	"time"

	"github.com/leeroyakbar/olap-backend/app/configs/clients"
)

type Dim_Time struct {
	Id_time  time.Time `json:"id" gorm:"primaryKey"`
	Tahun    int       `json:"tahun" form:"tahun"`
	Semester int       `json:"semester" form:"semester"`
	Kuartal  string    `json:"kuartal" form:"kuartal"`
	Bulan    string    `json:"bulan" form:"bulan"`
	Hari     string    `json:"hari" form:"hari"`
}

func (data Dim_Time) Add() error {
	return clients.DATABASE.Model(&data).Create(&data).Error
}

func (Dim_Time) TableName() string {
	return "dim_time"
}
