package customer

import (
	"github.com/ArWhale/WTest/internal/consts"
	"time"
)

type WebCustomer struct {
	ID        int64  `form:"id" json:"id,omitempty" binding:"-"`
	FirstName string `form:"firstName" json:"firstName" binding:"required"`
	LastName  string `form:"lastName" json:"lastName" binding:"required"`
	Gender    string `form:"gender" json:"gender" binding:"required,genderCustom"`
	Email     string `form:"email" json:"email" binding:"required"`
	Address   string `form:"address" json:"address" binding:"-"`
	Birthdate string `form:"birthdate" json:"birthdate" binding:"required,birthdateCustom"`
}

func (c *WebCustomer) ToDb() (*Customer, error) {
	dTime, err := time.Parse(consts.DefaultDateLayout, c.Birthdate)
	if err != nil {
		return nil, err
	}
	return &Customer{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Gender:    c.Gender,
		Email:     c.Email,
		Address:   c.Address,
		Birthdate: dTime,
	}, nil
}

type SearchCustomer struct {
	FirstName *string `form:"firstName" json:"firstName,omitempty"`
	LastName  *string `form:"lastName" json:"lastName,omitempty"`
	Limit     *int64  `form:"limit" json:"limit,omitempty"`
	Offset    *int64  `form:"offset" json:"offset,omitempty"`
}

type Customer struct {
	ID        int64     `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Gender    string    `json:"gender"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Birthdate time.Time `form:"birthdate" json:"birthdate" binding:"required,birthdateCustom"`
}

func (c *Customer) ToWeb() *WebCustomer {
	return &WebCustomer{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Gender:    c.Gender,
		Email:     c.Email,
		Address:   c.Address,
		Birthdate: c.Birthdate.Format(consts.DefaultDateLayout),
	}
}
