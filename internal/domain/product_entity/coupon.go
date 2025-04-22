package product_entity

import "time"

type CouponEntity struct {
	Id         string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	ProductId  string    `json:"product_id" gorm:"column:ProductId;unique" faker:"uuid_digit"`
	Quantity   int       `json:"quantity" gorm:"column:Quantity" faker:"boundary_start=1, boundary_end=1000"`
	Type       string    `json:"type" gorm:"column:Type" faker:"oneof: percentage, fixed"`
	Value      int64     `json:"value" gorm:"column:Value" faker:"boundary_start=1, boundary_end=100000"`
	ExpiryDate time.Time `json:"expiry_date" gorm:"column:ExpiryDate" faker:"time"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version    time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted  bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt  time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
	
	// Relationships
	Product    ProductEntity `json:"product" gorm:"foreignKey:ProductId"`
}

func (CouponEntity) TableName() string {
	return "Product.Coupons"
} 