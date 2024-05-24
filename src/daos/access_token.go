package daos

import (
	"log"
	"real_estate/src/database/models"
	"real_estate/src/utils/context"
)

type AccessTokenDAO interface {
	Create(ctx *context.Context, accessToken *models.AccessToken) error
	Upsert(ctx *context.Context, accessToken *models.AccessToken) error
	Get(ctx *context.Context, token string) (*models.AccessToken, error)
	Delete(ctx *context.Context, token string) error
}

func NewAccessToken() AccessTokenDAO {
	return &AccessToken{}
}

type AccessToken struct {
}

func (a *AccessToken) Create(ctx *context.Context, accessToken *models.AccessToken) error {
	err := ctx.DB.Table("access_tokens").Create(accessToken).Error
	if err != nil {
		log.Println("Unable to create AccessToken. Err:", err)
		return err
	}

	return nil
}

func (a *AccessToken) Upsert(ctx *context.Context, accessToken *models.AccessToken) error {
	err := ctx.DB.Table("access_tokens").Save(accessToken).Error
	if err != nil {
		log.Println("Unable to create AccessToken. Err:", err)
		return err
	}

	return nil
}

func (a *AccessToken) Get(ctx *context.Context, token string) (*models.AccessToken, error) {
	accesstoken := &models.AccessToken{}
	err := ctx.DB.Table("access_tokens").First(accesstoken, "token=?", token).Error
	if err != nil {
		log.Println("Unable to get AccessToken. Err:", err)
		return nil, err
	}

	return accesstoken, nil
}

func (a *AccessToken) Delete(ctx *context.Context, token string) error {
	err := ctx.DB.Table("access_tokens").Delete(&models.AccessToken{
		Token: token,
	}).Error
	if err != nil {
		log.Println("Unable to delete Token. Err:", err)
		return err
	}

	return nil
}
