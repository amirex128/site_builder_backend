package product_entity

type CategoryProductEntity struct {
	Id         string `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	ProductId  string `json:"product_id" gorm:"column:ProductId" faker:"uuid_digit"`
	CategoryId string `json:"category_id" gorm:"column:CategoryId" faker:"uuid_digit"`
	
	// Relationships
	Product    ProductEntity  `json:"product" gorm:"foreignKey:ProductId"`
	Category   CategoryEntity `json:"category" gorm:"foreignKey:CategoryId"`
}

func (CategoryProductEntity) TableName() string {
	return "Product.CategoryProduct"
} 