package categories

import (
	"real_estate/src/constants"
	"real_estate/src/daos"
	"real_estate/src/database/models"
	"real_estate/src/dtos"
	"real_estate/src/utils/context"

	"github.com/google/uuid"
)

type Category struct {
	category daos.CategoriesDAO
}

//TO access all daos
func New() *Category{
	return &Category{
		category: daos.NewCategories(),
	}
}   

//map dto to model
func(c *Category) categoryFromCategoryReq(req *dtos.CategoryReq) *models.Categories{
	return &models.Categories{
		Id: uuid.New().String(),
		Name: req.Name,
	}
}

func(c *Category) CreateCategory(ctx *context.Context,req *dtos.CategoryReq)error{
	category:=c.categoryFromCategoryReq(req)
	if ok,_:=c.category.CheckCategoryExist(ctx,category.Name);ok{
		return constants.ErrCategoryTaken
	}
	return c.category.Create(ctx,category)
}

