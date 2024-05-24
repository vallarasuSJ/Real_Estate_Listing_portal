package daos

import (
	"log"
	"real_estate/src/database/models"
	"real_estate/src/utils/context"
)

type PropertiesDAO interface {
	Create(ctx *context.Context, property *models.Properties) error
	Upsert(ctx *context.Context, property *models.Properties) error
	Get(ctx *context.Context, id string) (*models.Properties, error)
	Delete(ctx *context.Context, id string) error
	GetAll(ctx *context.Context) (*[]models.Properties, error)
	IsPropertyBooked(ctx *context.Context, id string) (bool, error)
}

func NewProperties() PropertiesDAO {
	return &Properties{}
}

type Properties struct {
}

func (p *Properties) Create(ctx *context.Context, property *models.Properties) error {
	err := ctx.DB.Table("properties").Create(property).Error
	if err != nil {
		log.Println("Unable to create property.Err:", err)
		return err
	}
	return nil
}

func (p *Properties) Upsert(ctx *context.Context, property *models.Properties) error {
	err := ctx.DB.Table("properties").Save(&property).Error
	if err != nil {
		log.Println("Unable to Update property.Err:", err)
		return err
	}
	return nil
}

func (p *Properties) Get(ctx *context.Context, id string) (*models.Properties, error) {
	property := &models.Properties{}
	err := ctx.DB.Table("properties").First(property, "id=?", id).Error
	if err != nil {
		log.Println("Unable to Get property.Err:", err)
		return nil, err
	}
	log.Println(property)
	return property, nil
}

func (p *Properties) GetAll(ctx *context.Context) (*[]models.Properties, error) {
	var properties *[]models.Properties
	var err error
	if ctx.Users.RoleName == "admin" {
		err = ctx.DB.Find(&properties).Error
		
	}else {
		err = ctx.DB.Where("is_approved = ? AND is_booked=?", true, false).Find(&properties).Error
	}
	if err != nil {
		return nil, err
	}
	return properties, nil
}

func (p *Properties) Delete(ctx *context.Context, id string) error {
	err := ctx.DB.Table("properties").Delete(&models.Properties{
		Id: id,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *Properties) IsPropertyBooked(ctx *context.Context, id string) (bool, error) {
	var property *models.Properties
	err := ctx.DB.Table("properties").Where("id=?", id).First(&property).Error
	if err!=nil{
		return false,err
	}
	return property.IsBooked, nil
}
