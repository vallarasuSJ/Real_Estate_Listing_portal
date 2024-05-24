package properties

import (
	"errors"
	"real_estate/src/constants"
	"real_estate/src/daos"
	"real_estate/src/database/models"
	"real_estate/src/dtos"
	properties "real_estate/src/services/Properties"
	"time"

	"real_estate/src/utils/context"

	"github.com/google/uuid"
)

type Bookings struct {
	booking  daos.Booked_propertiesDAO
	property properties.Property
}

func New() *Bookings {
	return &Bookings{
		booking:  daos.NewBooked_properties(),
		property: *properties.New(),
	}
}

func (b *Bookings) BookingFromBookingReq(req *dtos.BookingReq, userId string) *models.Booked_properties {
	return &models.Booked_properties{
		Id:         uuid.New().String(),
		UserId:     userId,
		PropertyId: req.Property,
		Created_at: time.Now(),
	}
}

func (b *Bookings) CreateBookings(ctx *context.Context, req *dtos.BookingReq) error {
	userId := ctx.Users.Id
	status,err:= b.property.CheckPropertyAlreadyBooked(ctx, req.Property)
	if err!=nil{
		return constants.ErrPropertyStatus
	}
	if !status{
		booking := b.BookingFromBookingReq(req, userId)
	    err = b.booking.Create(ctx, booking)
		if err!=nil{
			return constants.ErrBooking
		}
	}else{
		return errors.New("property already booked")
	}
	return  b.property.UpdateBookedProperty(ctx, req.Property) 
}


func (b *Bookings) GetBooking(ctx *context.Context) (*[]models.Booked_properties, error) {
	userId := ctx.Users.Id
	booking, err := b.booking.Get(ctx, userId)
	if err != nil {
		return nil, constants.ErrBookingNotExist
	}
	return booking, nil
}

func (b *Bookings) CancelBooking(ctx *context.Context, id string) error {
	var booking *models.Booked_properties
	booking, err := b.booking.Get_by_id(ctx, id)
	if err != nil {
		return constants.ErrBookingNotExist
	}
	err=b.property.UpdateBookedProperty(ctx, booking.PropertyId)
	if err != nil {
		return constants.ErrPropertyStatus
	}
	err = b.booking.Delete(ctx, id)
	if err != nil {
		return constants.ErrBookingNotExist
	}
	return nil
}
