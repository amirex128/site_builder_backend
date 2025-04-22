package product_entity

import "time"

type DiscountEntity struct {
	Id         string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	Code       string    `json:"code" gorm:"column:Code" faker:"word"`
	Quantity   int       `json:"quantity" gorm:"column:Quantity" faker:"boundary_start=1, boundary_end=1000"`
	Type       string    `json:"type" gorm:"column:Type" faker:"oneof: percentage, fixed"`
	Value      int64     `json:"value" gorm:"column:Value" faker:"boundary_start=1, boundary_end=100000"`
	ExpiryDate time.Time `json:"expiry_date" gorm:"column:ExpiryDate" faker:"time"`
	SiteId     string    `json:"site_id" gorm:"column:SiteId" faker:"uuid_digit"`
	UserId     string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version    time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted  bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt  time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
	
	// Relationships
	Products   []ProductEntity   `json:"products,omitempty" gorm:"many2many:Product.DiscountProduct;foreignKey:Id;joinForeignKey:DiscountId;References:Id;joinReferences:ProductId"`
	Customers  []CustomerDiscountEntity `json:"customers,omitempty" gorm:"foreignKey:DiscountId"`
}

func (DiscountEntity) TableName() string {
	return "Product.Discounts"
} 