package controllers

import (
	"net/http"
	"petlover/lib/repo"
	"petlover/models"
	"petlover/utils"

	"github.com/labstack/echo/v4"
)

func CreatePetController(c echo.Context) error {
	var petsRequest models.PetsRequest
	c.Bind(&petsRequest)

	if len(petsRequest.Name) == 0 {
		return c.JSON(http.StatusBadRequest, nil)
	}

	var petsDB = models.Pets{
		Name:   petsRequest.Name,
		Type: petsRequest.Type,
		Age: petsRequest.Age,}

	err := repo.InsertPet(&petsDB)
	if err != nil {
		return GenerateResponseError(c, err, nil)
	}

	var petsResponse = models.PetsResponse{
		Id:        petsDB.Id,
		Name:     petsDB.Name,
		Type:   petsDB.Type,
		Age: petsDB.Age,
		CreatedAt: petsDB.CreatedAt,
		UpdatedAt: petsDB.UpdatedAt,
	}

	return GenerateResponseSuccess(c, utils.SUCCESS_INSERT_DATA, petsResponse)
}

func GetPetController(c echo.Context) error {
	data, err := repo.GetPet()
	if err != nil {
		return GenerateResponseError(c, err, nil)
	}

	var petsResponse []models.PetsResponse
	for _, v := range data {
		petsResponse = append(petsResponse, models.PetsResponse{
			Id:        v.Id,
			Name:   v.Name,
			Type:     v.Type,
			Age:  v.Age,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return GenerateResponseSuccess(c, utils.SUCCESS_GET_DATA, petsResponse)
}
