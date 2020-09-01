package product

import (
	"github.com/gocql/gocql"
)

// Product struct to handle the product data to/from db
type Product struct {
	ID       gocql.UUID `json:"id"`
	Name     string     `json:"productname"`
	Descript string     `json:"productdescription"`
	Price    float64    `json:"productprice"`
	Active   bool       `json:"productactive"`
}

// GetProductResponse to form payload returning a single Product struct
type GetProductResponse struct {
	Product Product `json:"Product"`
}

// AllProductsResponse to form payload of an array of Product structs
type AllProductsResponse struct {
	Products []Product `json:"Products"`
}

// NewProductResponse builds a payload of new Product resource ID
type NewProductResponse struct {
	ID gocql.UUID `json:"id"`
}

// ErrorResponse returns an array of error strings if appropriate
type ErrorResponse struct {
	Errors []string `json:"errors"`
}
