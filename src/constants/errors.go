package constants

import "errors"

var (
	ErrEmailTaken  = errors.New("email taken")
	ErrMobileTaken = errors.New("mobile taken")
	ErrInValidCredentials=errors.New("!Invalid email or mobile or password")
	ErrRoleTaken=errors.New("role already exists")
	ErrAdminTaken=errors.New("cannot create account for the role admin") 
	ErrAccessTokenExpire = errors.New("access token is expired")
	ErrCategoryTaken=errors.New("category already exists")
	ErrPropertyNotExist=errors.New("unable to  find property") 
	ErrPropertyStatus=errors.New("unable to update the property status or property not approved") 
	ErrBookingNotExist=errors.New("cannot find the booking")
	ErrBooking=errors.New("unable to book property")
	ErrAccessTokenError = errors.New("not able to find user details")

)