package order_entity

import "time"

type OrderEntity struct {
	Id                           string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	SiteId                       string    `json:"site_id" gorm:"column:SiteId" faker:"uuid_digit"`
	TotalRawPrice                int64     `json:"total_raw_price" gorm:"column:TotalRawPrice" faker:"boundary_start=1000, boundary_end=1000000"`
	TotalCouponDiscount          int64     `json:"total_coupon_discount" gorm:"column:TotalCouponDiscount" faker:"boundary_start=0, boundary_end=100000"`
	TotalPriceWithCouponDiscount int64     `json:"total_price_with_coupon_discount" gorm:"column:TotalPriceWithCouponDiscount" faker:"boundary_start=1000, boundary_end=1000000"`
	CourierPrice                 int64     `json:"courier_price" gorm:"column:CourierPrice" faker:"boundary_start=0, boundary_end=50000"`
	Courier                      string    `json:"courier" gorm:"column:Courier" faker:"oneof: post, express, pickup"`
	OrderStatus                  string    `json:"order_status" gorm:"column:OrderStatus" faker:"oneof: pending, processing, shipped, delivered, cancelled"`
	TotalFinalPrice              int64     `json:"total_final_price" gorm:"column:TotalFinalPrice" faker:"boundary_start=1000, boundary_end=1000000"`
	Description                  string    `json:"description,omitempty" gorm:"column:Description" faker:"paragraph"`
	TotalWeight                  int       `json:"total_weight" gorm:"column:TotalWeight" faker:"boundary_start=100, boundary_end=10000"`
	TrackingCode                 string    `json:"tracking_code,omitempty" gorm:"column:TrackingCode" faker:"uuid_digit"`
	BasketId                     string    `json:"basket_id" gorm:"column:BasketId" faker:"uuid_digit"`
	DiscountId                   string    `json:"discount_id,omitempty" gorm:"column:DiscountId" faker:"uuid_digit"`
	AddressId                    string    `json:"address_id" gorm:"column:AddressId" faker:"uuid_digit"`
	CustomerId                   string    `json:"customer_id" gorm:"column:CustomerId" faker:"uuid_digit"`
	CreatedAt                    time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt                    time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version                      time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted                    bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt                    time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`

	// Relationships
	OrderItems []OrderItemEntity `json:"order_items,omitempty" gorm:"foreignKey:OrderId"`
}

func (OrderEntity) TableName() string {
	return "Order.Orders"
}
