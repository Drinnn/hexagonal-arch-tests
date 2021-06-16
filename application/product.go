package application

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetId() string
	GetName() string
	GetPrice() float64
	GetStatus() string
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Product struct {
	Id     string  `valid:"uuidv4"`
	Name   string  `valid:"required"`
	Price  float64 `valid:"float,optional"`
	Status string  `valid:"required"`
}

func (p *Product) IsValid() (bool, error) {
	if p.Status == "" {
		p.Status = DISABLED
	}

	if p.Status != ENABLED && p.Status != DISABLED {
		return false, errors.New("status must be ENABLED or DISABLED")
	}

	if p.Price < 0 {
		return false, errors.New("price must be greater than or equal 0")
	}

	_, err := govalidator.ValidateStruct(p)
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New("price must be greater than 0 to enable the product")
}

func (p *Product) Disable() error {
	if p.Price == 0 {
		p.Status = DISABLED
		return nil
	}

	return errors.New("price must be 0 to disable the product")
}

func (p *Product) GetId() string {
	return p.Id
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

func (p *Product) GetStatus() string {
	return p.Status
}
