package daos

import (
	"log"
	"real_estate/src/database/models"
	"real_estate/src/utils/context"
)

type Booked_propertiesDAO interface {
	Create(ctx *context.Context, booked_property *models.Booked_properties) error
	Upsert(ctx *context.Context, booked_property *models.Booked_properties) error
	Get(ctx *context.Context, id string) (*[]models.Booked_properties, error)
	Get_by_id(ctx *context.Context, id string) (*models.Booked_properties, error)
	Delete(ctx *context.Context, id string) error 

}

func NewBooked_properties() Booked_propertiesDAO {
	return &Booked_properties{}
}

type Booked_properties struct {
}


func (a *Booked_properties) Create(ctx *context.Context, booked_property *models.Booked_properties) error {
	err := ctx.DB.Table("booked_properties").Create(booked_property).Error
	if err != nil {
		log.Println("Unable to create booked_property.Err:", err)
		return err
	}
	return nil
}

func (a *Booked_properties) Upsert(ctx *context.Context, booked_property *models.Booked_properties) error {
	err := ctx.DB.Table("booked_properties").Save(booked_property).Error
	if err != nil {
		log.Println("Unable to Update booked_property.Err:", err)
		return err
	}
	return nil
}

func (a *Booked_properties) Get(ctx *context.Context, id string) (*[]models.Booked_properties, error) {
	booked_properties := &[]models.Booked_properties{} 
	var err error
	if ctx.Users.RoleName=="admin"{
		err = ctx.DB.Find(&booked_properties).Error

	}else{
		err = ctx.DB.Where("user_id=?", id).Find(&booked_properties).Error
	}
	if err != nil {
		log.Println("Unable to Get booked_property.Err:", err)
		return nil, err
	}
	return booked_properties, nil
}

func (a *Booked_properties) Get_by_id(ctx *context.Context, id string) (*models.Booked_properties, error) {
	booked_properties := &models.Booked_properties{} 
	err := ctx.DB.Table("booked_properties").Where("id=?", id).First(&booked_properties).Error
	if err != nil {
		log.Println("Unable to Get booked_property.Err:", err)
		return nil, err
	}
	return booked_properties, nil
}


func (a *Booked_properties) Delete(ctx *context.Context, id string) error {
	err := ctx.DB.Table("booked_properties").Delete(&models.Booked_properties{
		Id: id,
	}).Error
	if err != nil {
		log.Println("Unable to Delete booked_property.Err:", err)
		return err
	}
	return nil
}
