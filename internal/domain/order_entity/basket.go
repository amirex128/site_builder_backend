package order_entity

import "time"

type BasketEntity struct {
	Id                           string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	SiteId                       string    `json:"site_id" gorm:"column:SiteId" faker:"uuid_digit"`
	TotalRawPrice                int64     `json:"total_raw_price" gorm:"column:TotalRawPrice" faker:"boundary_start=1000, boundary_end=1000000"`
	TotalCouponDiscount          int64     `json:"total_coupon_discount" gorm:"column:TotalCouponDiscount" faker:"boundary_start=0, boundary_end=100000"`
	TotalPriceWithCouponDiscount int64     `json:"total_price_with_coupon_discount" gorm:"column:TotalPriceWithCouponDiscount" faker:"boundary_start=1000, boundary_end=1000000"`
	DiscountId                   string    `json:"discount_id,omitempty" gorm:"column:DiscountId" faker:"uuid_digit"`
	CustomerId                   string    `json:"customer_id" gorm:"column:CustomerId" faker:"uuid_digit"`
	CreatedAt                    time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt                    time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version                      time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted                    bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt                    time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
	
	// Relationships
	BasketItems                  []BasketItemEntity `json:"basket_items,omitempty" gorm:"foreignKey:BasketId"`
}

func (BasketEntity) TableName() string {
	return "Order.Baskets"
} 