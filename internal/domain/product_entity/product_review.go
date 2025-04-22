package product_entity

import "time"

type ProductReviewEntity struct {
	Id         string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	Rating     int       `json:"rating" gorm:"column:Rating" faker:"boundary_start=1, boundary_end=5"`
	Like       int       `json:"like" gorm:"column:Like" faker:"boundary_start=0, boundary_end=100"`
	Dislike    int       `json:"dislike" gorm:"column:Dislike" faker:"boundary_start=0, boundary_end=100"`
	Approved   bool      `json:"approved" gorm:"column:Approved" faker:"oneof: true, false"`
	ReviewText string    `json:"review_text" gorm:"column:ReviewText" faker:"paragraph"`
	ProductId  string    `json:"product_id" gorm:"column:ProductId" faker:"uuid_digit"`
	SiteId     string    `json:"site_id" gorm:"column:SiteId" faker:"uuid_digit"`
	UserId     string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CustomerId string    `json:"customer_id" gorm:"column:CustomerId" faker:"uuid_digit"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version    time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted  bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt  time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`

	// Relationships
	Product ProductEntity `json:"product" gorm:"foreignKey:ProductId"`
}

func (ProductReviewEntity) TableName() string {
	return "Product.ProductReviews"
}
