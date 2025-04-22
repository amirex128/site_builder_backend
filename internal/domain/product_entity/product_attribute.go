package product_entity

import "time"

type ProductAttributeEntity struct {
	Id        string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	ProductId string    `json:"product_id" gorm:"column:ProductId" faker:"uuid_digit"`
	Type      string    `json:"type" gorm:"column:Type" faker:"oneof: color, size, material, weight"`
	Name      string    `json:"name" gorm:"column:Name" faker:"word"`
	Value     string    `json:"value" gorm:"column:Value" faker:"word"`
	CreatedAt time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version   time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
	
	// Relationships
	Product   ProductEntity `json:"product" gorm:"foreignKey:ProductId"`
}

func (ProductAttributeEntity) TableName() string {
	return "Product.ProductAttributes"
} 