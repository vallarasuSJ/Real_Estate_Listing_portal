package users

import (
	"log"
	"real_estate/src/config"
	"real_estate/src/constants"
	"real_estate/src/daos"
	"real_estate/src/database/models"
	"real_estate/src/dtos"
	"real_estate/src/utils/context"
	"real_estate/src/utils/token"

	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	user daos.UsersDAO
	accessToken  daos.AccessTokenDAO
	refreshToken daos.RefreshTokenDAO
	role daos.RolesDAO
}

func New() *User{
	return &User{
		user: daos.NewUsers(),
		accessToken:  daos.NewAccessToken(),
		refreshToken: daos.NewRefreshToken(),
		role: daos.NewRole(),
	}
} 


func (a *User) SetDAOs(account daos.UsersDAO, accessToken daos.AccessTokenDAO,refreshToken daos.RefreshTokenDAO,role daos.RolesDAO){
	a.user = account
	a.accessToken = accessToken
	a.refreshToken = refreshToken
	a.role= role

}
func (a *User) GetAccountWithAccessToken(ctx *context.Context, token string) (*dtos.Users, error) {
	at, err := a.accessToken.Get(ctx, token)
	if err != nil {
		log.Println("Unable to get access token. Err: ", err)
		return nil, err
	}

	if at.ExpiresAt.Before(time.Now()) {
		return nil, constants.ErrAccessTokenExpire
	}

	account, err := a.user.Get(ctx, at.UserId)
	if err != nil {
		log.Println("Unable to get account. Err: ", err)
		return nil, err
	}
	return account, nil
}

func (a *User) GetAccessFromRefreshToken(ctx *context.Context, tkn string) (string, error) {
	rt, err := a.refreshToken.Get(ctx, tkn)
	if err != nil {
		log.Println("Unable to get access token. Err: ", err)
		return "", err
	}

	if rt.ExpiresAt.Before(time.Now()) {
		return "", constants.ErrAccessTokenExpire
	}

	accessToken, _ := token.GetAccessAndRefreshToken(config.Conf.TokenSize)
	err = a.accessToken.Create(ctx, &models.AccessToken{
		Token:         accessToken,
		UserId:     rt.UserId,
		RefreshTokens: rt.Token,
		ExpiresAt:     time.Now().Add(time.Duration(config.Conf.AccessTokenExpiry) * time.Hour),
	})

	if err != nil {
		log.Println("Unable to get access token. Err: ", err)
		return "", err
	}
	return accessToken, nil
}


func( a *User) accountFromRegisterReq(ctx *context.Context,req *dtos.RegisterReq) *models.Users{
	role,err:= a.role.GetRoleByName(ctx,req.Role) 
	if err!=nil{
		log.Println("Unable to get role,Err: ",err)
	}
	return &models.Users{
		Id: uuid.New().String(),
		Email: req.Email,
		Contact_number: req.Contact_Number,
		Password: req.Password,
		Username: req.Username,
		Gender:req.Gender,
		Role_id:role,
		Created_at: time.Now(),
	}
}


func( a *User) Register(ctx *context.Context,req *dtos.RegisterReq) error{  
	if(req.Role=="admin"){
		return constants.ErrAdminTaken
	} 
	user:= a.accountFromRegisterReq(ctx,req) 

	hash,err:=bcrypt.GenerateFromPassword([]byte(user.Password),10) 
	if err!=nil{
		log.Println("Unable to create password hash,Err: ",err)
		return err
	}

	if ok,_:=a.user.CheckEmailExists(ctx,user.Email,user.Role_id);ok{
		return constants.ErrEmailTaken
	}
	
	if ok,_:=a.user.CheckMobileExists(ctx,user.Contact_number,user.Role_id);ok{
		return constants.ErrMobileTaken
	}


	user.Password=string(hash) 

	return a.user.Create(ctx,user)

}

func( a *User) Login(ctx *context.Context,req *dtos.LoginReq) (*dtos.LoginRes,error){ 
	user,err:=a.user.GetAccountForEmailorMobile(ctx,req.Email,req.Contact_number) 
	if err==gorm.ErrRecordNotFound{
		log.Println("No record found.Err: ",err)
		return nil,err
	}
	if err!=nil{
		log.Println("Incorrect email or mobile or password",err)
		return nil,err
	} 


	err=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(req.Password)) 
	if err!=nil{
		log.Println("Password is incorrect,Err: ",err)
		return nil,err
	} 

	accessToken, refreshToken := token.GetAccessAndRefreshToken(config.Conf.TokenSize)

	a.refreshToken.Create(ctx, &models.RefreshToken{
		Token:     refreshToken,
		UserId: user.Id,
		ExpiresAt: time.Now().Add(time.Duration(config.Conf.RefreshTokenExpiry) * time.Hour),
	})
	a.accessToken.Create(ctx, &models.AccessToken{
		Token:        accessToken,
		RefreshTokens: refreshToken,
		UserId:    user.Id,
		ExpiresAt:    time.Now().Add(time.Duration(config.Conf.AccessTokenExpiry) * time.Hour),
	})

	return &dtos.LoginRes{
		Token: accessToken,
	},nil

}