package order_entity

import "time"

type ReturnItemEntity struct {
	Id           string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	ReturnReason string    `json:"return_reason" gorm:"column:ReturnReason" faker:"sentence"`
	OrderStatus  int       `json:"order_status" gorm:"column:OrderStatus" faker:"oneof: 1, 2, 3, 4, 5"`
	OrderItemId  string    `json:"order_item_id" gorm:"column:OrderItemId;unique" faker:"uuid_digit"`
	ProductId    string    `json:"product_id" gorm:"column:ProductId" faker:"uuid_digit"`
	UserId       string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CustomerId   string    `json:"customer_id" gorm:"column:CustomerId" faker:"uuid_digit"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version      time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted    bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt    time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
	
	// Relationships
	OrderItem    OrderItemEntity `json:"order_item" gorm:"foreignKey:OrderItemId"`
}

func (ReturnItemEntity) TableName() string {
	return "Order.ReturnItem"
} 