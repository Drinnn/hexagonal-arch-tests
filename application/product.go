package application

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

func (p *Product) IsValid() (bool, error) {

}

func (p *Product) Enable() error {

}

func (p *Product) Disable() error {

}

func (p *Product) GetId() string {

}

func (p *Product) GetName() string {

}

func (p *Product) GetPrice() float64 {

}

func (p *Product) GetStatus() string {

}