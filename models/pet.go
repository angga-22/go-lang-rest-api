package models

import (
	"time"

	"gorm.io/gorm"
)

// struct database
type Pets struct {
	Id        int            `json:"id"`
	Name     string         `json:"name"`
	Type		 string         `json:"type"`
	Age   	int         		`json:"age"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"deletedAt"`
}

type PetsRequest struct {
	Name   string `json:"name"`
	Type 	string `json:"type"`
	Age 	int `json:"age"`
}

type PetsResponse struct {
	Id        int       `json:"id"`
	Name     string    `json:"name"`
	Type   string    `json:"type"`
	Age 		int    `json:"age"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
