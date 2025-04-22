package product_entity

import "time"

type ProductVariantEntity struct {
	Id         string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	ProductId  string    `json:"product_id" gorm:"column:ProductId" faker:"uuid_digit"`
	Name       string    `json:"name" gorm:"column:Name" faker:"word"`
	Price      int64     `json:"price" gorm:"column:Price" faker:"boundary_start=1000, boundary_end=1000000"`
	Stock      int       `json:"stock" gorm:"column:Stock" faker:"boundary_start=0, boundary_end=1000"`
	UserId     string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CustomerId string    `json:"customer_id" gorm:"column:CustomerId" faker:"uuid_digit"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version    time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted  bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt  time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`

	// Relationships
	Product ProductEntity `json:"product" gorm:"foreignKey:ProductId"`
}

func (ProductVariantEntity) TableName() string {
	return "Product.ProductVariants"
}
