package roles
 
import (
    "real_estate/src/constants"
    "real_estate/src/daos"
    "real_estate/src/database/models"
    "real_estate/src/dtos"
    "real_estate/src/utils/context"
 
    "github.com/google/uuid"
)
 
type Role struct {
    role daos.RolesDAO
}
 
func New() *Role {
    return &Role{
        role: daos.NewRole(),
    }
}
 
func (r *Role) roleFromRegisterReq(req *dtos.RoleReq) *models.Roles {
    return &models.Roles{
        ID:   uuid.New().String(),
        Name: req.Name,
    }
}
 
func (r *Role) RegisterRoles(ctx *context.Context, req *dtos.RoleReq) error {
    role := r.roleFromRegisterReq(req)
    if ok, _ := r.role.CheckRoleExist(ctx, role.Name); ok {
        return constants.ErrRoleTaken
    }
    return r.role.Create(ctx, role)
}
 