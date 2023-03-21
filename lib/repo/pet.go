package repo

import (
	"petlover/config"
	"petlover/models"
	"petlover/utils"

	"gorm.io/gorm"
)

func InsertPet(pet *models.Pets) error {
	result := config.DB.Create(pet)
	if result.Error != nil {
		return utils.ERR_CREATE_DATABASE
	}
	return nil
}

func GetPet() ([]models.Pets, error) {
	var data []models.Pets
	result := config.DB.Find(&data)

	if result.Error == gorm.ErrRecordNotFound {
		return []models.Pets{}, utils.ERR_EMPTY_DATA_IN_DATABASE
	} else if result.Error != nil {
		return nil, utils.ERR_GET_DATA_IN_DATABASE
	}

	return data, nil
}
