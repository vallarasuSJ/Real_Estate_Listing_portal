package daos

import (
	"log"
	"real_estate/src/database/models"
	"real_estate/src/dtos"
	"real_estate/src/utils/context"
)

type UsersDAO interface {
	Create(ctx *context.Context, user *models.Users) error
	Upsert(ctx *context.Context, user *models.Users) error
	Get(ctx *context.Context, id string) (*dtos.Users, error)
	Delete(ctx *context.Context, id string) error
	CheckEmailExists(ctx *context.Context,email,role string)(bool,error) 
	CheckMobileExists(ctx *context.Context,mobile,role string)(bool,error)
	GetAccountForEmailorMobile(ctx *context.Context, email, mobile string) (*models.Users, error)
}

func NewUsers() UsersDAO {
	return &Users{}
}

type Users struct {
}

func (a *Users) GetAccountForEmailorMobile(ctx *context.Context, email, contact_number string) (*models.Users, error) {
	user := &models.Users{}
	err := ctx.DB.Table("users").Select("*").Where("email=? AND contact_number=?", email, contact_number).First(user).Error
	if err != nil {
		log.Println("Unable to fetch login details.Err: ", err)
		return nil, err
	}
	return user, nil
}


func (a *Users) CheckEmailExists(ctx *context.Context,email,role string)(bool,error){ 
	var cnt int
	err := ctx.DB.Table("users").Select("count(*)").Where("email=? AND role_id=?",email,role).Scan(&cnt).Error
	if err != nil {
		log.Println("Unable to create user.Err:", err)
		return false,err
	}
	return cnt>0,nil
}

func (a *Users) CheckMobileExists(ctx *context.Context,mobile,role string)(bool,error){ 
	var cnt int
	err := ctx.DB.Table("users").Select("count(*)").Where("contact_number=? AND role_id=?",mobile,role).Scan(&cnt).Error
	if err != nil {
		log.Println("Unable to find user.Err:", err)
		return false,err
	}
	return cnt>0,nil
}

func (a *Users) Create(ctx *context.Context, user *models.Users) error {
	err := ctx.DB.Table("users").Create(user).Error
	if err != nil {
		log.Println("Unable to create user.Err:", err)
		return err
	}
	return nil
}

func (a *Users) Upsert(ctx *context.Context, user *models.Users) error {
	err := ctx.DB.Table("users").Save(user).Error
	if err != nil {
		log.Println("Unable to update user.Err:", err)
		return err
	}
	return nil
}

func (a *Users) Get(ctx *context.Context, id string) (*dtos.Users, error) {
	user := &dtos.Users{}
	err := ctx.DB.Table("users").
        Select("users.*, roles.name as role_name").
        Joins("inner join roles on roles.id = users.role_id").
        Where("users.id = ?", id).
        First(user).Error
	if err != nil {
		log.Println("Unable to get user.Err:", err)
		return nil, err
	}
	log.Println(user)
	return user, nil
}

func (a *Users) Delete(ctx *context.Context, id string) error {
	err := ctx.DB.Table("users").Delete(&models.Users{
		Id: id,
	}).Error
	if err != nil {
		log.Println("Unable to delete user.Err:", err)
		return err
	}
	return nil
}
