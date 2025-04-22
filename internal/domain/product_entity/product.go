package product_entity

import (
	"site_builder_backend/internal/domain/site_entity"
	"time"
)

type ProductEntity struct {
	Id              string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	Name            string    `json:"name" gorm:"column:Name" faker:"name"`
	Description     string    `json:"description,omitempty" gorm:"column:Description" faker:"paragraph"`
	Status          string    `json:"status" gorm:"column:Status" faker:"oneof: active, inactive, out_of_stock"`
	Weight          int       `json:"weight" gorm:"column:Weight" faker:"boundary_start=100, boundary_end=10000"`
	SellingCount    int       `json:"selling_count" gorm:"column:SellingCount" faker:"boundary_start=0, boundary_end=1000"`
	VisitedCount    int       `json:"visited_count" gorm:"column:VisitedCount" faker:"boundary_start=0, boundary_end=10000"`
	ReviewCount     int       `json:"review_count" gorm:"column:ReviewCount" faker:"boundary_start=0, boundary_end=500"`
	Rate            int       `json:"rate" gorm:"column:Rate" faker:"boundary_start=0, boundary_end=5"`
	Badges          string    `json:"badges,omitempty" gorm:"column:Badges" faker:"sentence"`
	FreeSend        bool      `json:"free_send" gorm:"column:FreeSend" faker:"oneof: true, false"`
	LongDescription string    `json:"long_description,omitempty" gorm:"column:LongDescription" faker:"paragraph"`
	Slug            string    `json:"slug" gorm:"column:Slug" faker:"slug"`
	SeoTags         string    `json:"seo_tags,omitempty" gorm:"column:SeoTags" faker:"sentence"`
	SiteId          string    `json:"site_id" gorm:"column:SiteId" faker:"uuid_digit"`
	UserId          string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CreatedAt       time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version         time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted       bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt       time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`

	// Relationships
	Categories []CategoryEntity                     `json:"categories,omitempty" gorm:"many2many:Product.CategoryProduct;foreignKey:Id;joinForeignKey:ProductId;References:Id;joinReferences:CategoryId"`
	Coupon     *CouponEntity                        `json:"coupon,omitempty" gorm:"foreignKey:ProductId"`
	Discounts  []DiscountEntity                     `json:"discounts,omitempty" gorm:"many2many:Product.DiscountProduct;foreignKey:Id;joinForeignKey:ProductId;References:Id;joinReferences:DiscountId"`
	Attributes []ProductAttributeEntity             `json:"attributes,omitempty" gorm:"foreignKey:ProductId"`
	Media      []ProductMediaEntity                 `json:"media,omitempty" gorm:"foreignKey:ProductId"`
	Variants   []ProductVariantEntity               `json:"variants,omitempty" gorm:"foreignKey:ProductId"`
	Reviews    []ProductReviewEntity                `json:"reviews,omitempty" gorm:"foreignKey:ProductId"`
	PageUsages []site_entity.PageProductUsageEntity `json:"page_usages,omitempty" gorm:"foreignKey:ProductId"`
}

func (ProductEntity) TableName() string {
	return "Product.Products"
}
