package application

import "errors"

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

type Product struct {
	Id     string
	Name   string
	Price  float64
	Status string
}

// func (p *Product) IsValid() (bool, error) {

// }

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}

	return errors.New("price must be greater than 0 to enable the product")
}

// func (p *Product) Disable() error {

// }

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
