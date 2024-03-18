package productdomain

import "time"

// const for error messages
const (
	ErrorSkuExists     string = "ERROR_SKU_EXISTS"
	ErrorWrongCategory string = "ERROR_WRONG_CATEGORY"
)

// domain representation of a product
type ProductDomain struct {
	Id        int64
	Name      string
	Sku       string
	Category  string
	Price     float64
	CreatedAt time.Time
}

// create a new product pointer instance
func New(n string, s string, c string, p float64) (*ProductDomain, error) {
	return &ProductDomain{
		Name:     n,
		Sku:      s,
		Category: c,
		Price:    p,
	}, nil
}
