package daos

import (
	"log"
	"real_estate/src/database/models"
	"real_estate/src/utils/context"
)
 
type RolesDAO interface {
    Create(ctx *context.Context, role *models.Roles) error
    CheckRoleExist(ctx*context.Context,name string)(bool,error)
    GetRoleByName(ctx *context.Context, roleName string) (string, error)
}
 
type Roles struct {
}
 
func NewRole() RolesDAO {
    return &Roles{}
}

func(r *Roles)GetRoleByName(ctx *context.Context, roleName string) (string, error) {
    var role models.Roles
    if err := ctx.DB.Where("name = ?", roleName).First(&role).Error; err != nil {
        return "", err
    }
    return role.ID, nil
}
 
func (r *Roles) Create(ctx *context.Context, role *models.Roles) error {
    err := ctx.DB.Table("roles").Create(role).Error
    if err != nil {
        log.Println("Unable to create Role. Err:", err)
        return err
    }
    return nil
}
 
func(r *Roles)CheckRoleExist(ctx*context.Context,name string)(bool,error){
    var cnt int
    err:=ctx.DB.Table("roles").Select("count(*)").Where("name=?",name).Scan(&cnt).Error
    if err != nil {
        log.Println("Unable to find role.Err:", err)
        return false,err
    }
    return cnt>0,nil
}