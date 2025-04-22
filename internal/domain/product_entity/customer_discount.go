package product_entity

type CustomerDiscountEntity struct {
	Id         string `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	DiscountId string `json:"discount_id" gorm:"column:DiscountId" faker:"uuid_digit"`
	CustomerId string `json:"customer_id" gorm:"column:CustomerId" faker:"uuid_digit"`
	
	// Relationships
	Discount   DiscountEntity `json:"discount" gorm:"foreignKey:DiscountId"`
}

func (CustomerDiscountEntity) TableName() string {
	return "Product.CustomerDiscount"
} 