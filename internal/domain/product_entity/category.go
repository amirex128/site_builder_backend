package product_entity

import (
	"time"
)

type CategoryEntity struct {
	Id               string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	Name             string    `json:"name" gorm:"column:Name" faker:"word"`
	ParentCategoryId string    `json:"parent_category_id,omitempty" gorm:"column:ParentCategoryId" faker:"uuid_digit"`
	Order            int       `json:"order" gorm:"column:Order" faker:"boundary_start=1, boundary_end=100"`
	Description      string    `json:"description,omitempty" gorm:"column:Description" faker:"paragraph"`
	Slug             string    `json:"slug" gorm:"column:Slug" faker:"slug"`
	SeoTags          string    `json:"seo_tags,omitempty" gorm:"column:SeoTags" faker:"sentence"`
	SiteId           string    `json:"site_id" gorm:"column:SiteId" faker:"uuid_digit"`
	CategoryId       string    `json:"category_id,omitempty" gorm:"column:CategoryId" faker:"uuid_digit"`
	UserId           string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CreatedAt        time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version          time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted        bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt        time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
	
	// Relationships
	ParentCategory   *CategoryEntity  `json:"parent_category,omitempty" gorm:"foreignKey:CategoryId"`
	ChildCategories  []CategoryEntity `json:"child_categories,omitempty" gorm:"foreignKey:CategoryId"`
	Products         []ProductEntity  `json:"products,omitempty" gorm:"many2many:Product.CategoryProduct;foreignKey:Id;joinForeignKey:CategoryId;References:Id;joinReferences:ProductId"`
	Media            []CategoryMediaEntity `json:"media,omitempty" gorm:"foreignKey:CategoryId"`
}

func (CategoryEntity) TableName() string {
	return "Product.Categories"
} 