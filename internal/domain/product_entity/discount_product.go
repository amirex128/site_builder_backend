package product_entity

type DiscountProductEntity struct {
	Id         string `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	ProductId  string `json:"product_id" gorm:"column:ProductId" faker:"uuid_digit"`
	DiscountId string `json:"discount_id" gorm:"column:DiscountId" faker:"uuid_digit"`
	
	// Relationships
	Product    ProductEntity  `json:"product" gorm:"foreignKey:ProductId"`
	Discount   DiscountEntity `json:"discount" gorm:"foreignKey:DiscountId"`
}

func (DiscountProductEntity) TableName() string {
	return "Product.DiscountProduct"
} 