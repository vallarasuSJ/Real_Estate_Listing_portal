package daos

import (
	"log"
	"real_estate/src/database/models"
	"real_estate/src/utils/context"
)

type CategoriesDAO interface {
	Create(ctx *context.Context, category *models.Categories) error
	Upsert(ctx *context.Context, category *models.Categories) error
	Get(ctx *context.Context, id string) (*models.Categories, error)
	Delete(ctx *context.Context, id string) error 
	CheckCategoryExist(ctx *context.Context,name string) (bool,error)
	
}

func NewCategories() CategoriesDAO {
	return &Categories{}
}

type Categories struct {
}


func (c *Categories) Create(ctx *context.Context, category *models.Categories) error {
	err := ctx.DB.Table("categories").Create(category).Error
	if err != nil {
		log.Println("Unable to create category.Err:", err)
		return err
	}
	return nil
}

func (c *Categories) Upsert(ctx *context.Context, category *models.Categories) error {
	err := ctx.DB.Table("categories").Save(category).Error
	if err != nil {
		log.Println("Unable to update category.Err:", err)
		return err
	}
	return nil
}

func (c *Categories) Get(ctx *context.Context, id string) (*models.Categories, error) {
	category := &models.Categories{}
	err := ctx.DB.Table("categories").First(category, "id=?", id).Error
	if err != nil {
		log.Println("Unable to Get category.Err:", err)
		return nil, err
	}
	return category, nil
}

func (c *Categories) Delete(ctx *context.Context, id string) error {
	err := ctx.DB.Table("categories").Delete(&models.Categories{
		Id: id,
	}).Error
	if err != nil {
		log.Println("Unable to delete category.Err:", err)
		return err
	}
	return nil
}

func(c *Categories)CheckCategoryExist(ctx *context.Context,name string)(bool,error){ 
	var count int
	err:=ctx.DB.Table("categories").Select("count(*)").Where("name=?",name).Scan(&count).Error
	if err!=nil{
		log.Println("Unable to find category.Err:", err)
        return false, err
	} 
	return count>0,nil
}
