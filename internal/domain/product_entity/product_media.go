package product_entity

type ProductMediaEntity struct {
	Id        string `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	ProductId string `json:"product_id" gorm:"column:ProductId" faker:"uuid_digit"`
	MediaId   string `json:"media_id" gorm:"column:MediaId" faker:"uuid_digit"`
	
	// Relationships
	Product   ProductEntity `json:"product" gorm:"foreignKey:ProductId"`
}

func (ProductMediaEntity) TableName() string {
	return "Product.ProductMedia"
} 