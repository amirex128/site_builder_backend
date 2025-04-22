package order_entity

import "time"

type BasketItemEntity struct {
	Id                           string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	Quantity                     int       `json:"quantity" gorm:"column:Quantity" faker:"boundary_start=1, boundary_end=10"`
	RawPrice                     int64     `json:"raw_price" gorm:"column:RawPrice" faker:"boundary_start=1000, boundary_end=100000"`
	FinalRawPrice                int64     `json:"final_raw_price" gorm:"column:FinalRawPrice" faker:"boundary_start=1000, boundary_end=100000"`
	FinalPriceWithCouponDiscount int64     `json:"final_price_with_coupon_discount" gorm:"column:FinalPriceWithCouponDiscount" faker:"boundary_start=1000, boundary_end=100000"`
	JustCouponPrice              int64     `json:"just_coupon_price" gorm:"column:JustCouponPrice" faker:"boundary_start=0, boundary_end=10000"`
	JustDiscountPrice            int64     `json:"just_discount_price" gorm:"column:JustDiscountPrice" faker:"boundary_start=0, boundary_end=10000"`
	BasketId                     string    `json:"basket_id" gorm:"column:BasketId" faker:"uuid_digit"`
	ProductId                    string    `json:"product_id" gorm:"column:ProductId" faker:"uuid_digit"`
	ProductVariantId             string    `json:"product_variant_id" gorm:"column:ProductVariantId" faker:"uuid_digit"`
	CreatedAt                    time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt                    time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version                      time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted                    bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt                    time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
	
	// Relationships
	Basket                       BasketEntity `json:"basket" gorm:"foreignKey:BasketId"`
}

func (BasketItemEntity) TableName() string {
	return "Order.BasketItems"
} 