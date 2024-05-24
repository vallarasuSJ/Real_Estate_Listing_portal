package properties

import (
	"real_estate/src/constants"
	"real_estate/src/daos"
	"real_estate/src/database/models"
	"real_estate/src/dtos"
	"real_estate/src/utils/context"
	"time"
	"github.com/google/uuid"
)

type Property struct {
	property daos.PropertiesDAO
}

func New() *Property {
	return &Property{
		property: daos.NewProperties(),
	}
}

func (p *Property) PropertyFromPropertyReq(req *dtos.PropertyReq, userId string) *models.Properties {
	return &models.Properties{
		Id:         uuid.New().String(),
		Name:       req.Name,
		Location:   req.Location,
		Price:      req.Price,
		UserId:     userId,
		CategoryId: req.Category, 
		Created_at: time.Now(),
	}
}

func (p *Property) CreateProperty(ctx *context.Context, req *dtos.PropertyReq) error {
	userId:=ctx.Users.Id
	property := p.PropertyFromPropertyReq(req, userId)

	return p.property.Create(ctx, property)
}

func (p *Property) GetAllProperties(ctx *context.Context) (*[]models.Properties, error) {
	properties, err := p.property.GetAll(ctx) 
	if err != nil {
		return nil, constants.ErrPropertyNotExist
	}
	return properties, nil
}

func (p *Property) GetProperty(ctx *context.Context, id string) (*models.Properties, error) {
	property, err := p.property.Get(ctx, id)
	if err != nil {
		return nil, constants.ErrPropertyNotExist
	}
	return property, nil
}

func (p *Property) UpdateProperty(ctx *context.Context, req *dtos.PropertyReq, id string) error {
	userId:=ctx.Users.Id

	propertyDetails, err := p.property.Get(ctx, id)
	if err != nil {
		return constants.ErrPropertyNotExist
	}
	property := &models.Properties{
		Id:         propertyDetails.Id,
		Name:       req.Name,
		Location:   req.Location,
		Price:      req.Price,
		UserId:     userId,
		CategoryId: propertyDetails.CategoryId,
	}
	return p.property.Upsert(ctx, property)
}

func (p *Property) ApproveProperty(ctx *context.Context, id string) error {
	propertyDetails, err := p.property.Get(ctx, id)
	if err != nil {
		return constants.ErrPropertyNotExist
	}
	property :=propertyDetails
	property.IsApproved=true
	
	return p.property.Upsert(ctx, property)
}

func (p *Property) UpdateBookedProperty(ctx *context.Context, id string) error { 
	propertyDetails, err := p.property.Get(ctx, id)
	if err != nil {
		return constants.ErrPropertyNotExist
	}
	property :=propertyDetails 
	if property.IsBooked{
		property.IsBooked=false
	}else{
		property.IsBooked=true
	}
	return p.property.Upsert(ctx, property)
}

func (p *Property) DeleteProperty(ctx *context.Context, id string) error {
	err := p.property.Delete(ctx, id)
	if err != nil {
		return constants.ErrPropertyNotExist
	}
	return nil
} 

func(p *Property) CheckPropertyAlreadyBooked(ctx *context.Context,id string) (bool,error){
	status,err:=p.property.IsPropertyBooked(ctx,id)
	if err!=nil{
		return false,constants.ErrPropertyNotExist
	} 
	if status==true{
		return false,constants.ErrPropertyStatus
	}
	return status,nil
}
