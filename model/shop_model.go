package model

type Owner struct {
	ID    uint
	Name  string
	Email string
	Phone string
	Shops []Shop `gorm:"foreignKey:OwnerID"` // One-to-many relationship with Shops
}

// Shop represents the Shops table
type Shop struct {
	ID       uint
	Name     string
	Address  string
	OwnerID  uint      // Foreign key referencing Owners table
	Owner    Owner     // Owner relation
	Products []Product `gorm:"foreignKey:ShopID"` // One-to-many relationship with Products
}
