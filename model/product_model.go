package model

type Product struct {
	ID     uint
	Name   string
	Price  float64
	ShopID uint // Foreign key referencing Shops table
	Shop   Shop // Shop relation
}
