package site_entity

import "time"

type PageEntity struct {
	Id          string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	SiteId      string    `json:"site_id" gorm:"column:SiteId" faker:"uuid_digit"`
	HeaderId    string    `json:"header_id" gorm:"column:HeaderId" faker:"uuid_digit"`
	FooterId    string    `json:"footer_id" gorm:"column:FooterId" faker:"uuid_digit"`
	Slug        string    `json:"slug" gorm:"column:Slug" faker:"slug"`
	Title       string    `json:"title" gorm:"column:Title" faker:"sentence"`
	Description string    `json:"description,omitempty" gorm:"column:Description" faker:"paragraph"`
	Body        string    `json:"body,omitempty" gorm:"column:Body" faker:"paragraph"`
	SeoTags     string    `json:"seo_tags,omitempty" gorm:"column:SeoTags" faker:"sentence"`
	UserId      string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version     time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted   bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt   time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
	
	// Relationships
	Site        SiteEntity       `json:"site" gorm:"foreignKey:SiteId"`
	Header      HeaderFooterEntity `json:"header" gorm:"foreignKey:HeaderId"`
	Footer      HeaderFooterEntity `json:"footer" gorm:"foreignKey:FooterId"`
	Media       []PageMediaEntity  `json:"media,omitempty" gorm:"foreignKey:PageId"`
	ArticleUsages []PageArticleUsageEntity `json:"article_usages,omitempty" gorm:"foreignKey:PageId"`
	ProductUsages []PageProductUsageEntity `json:"product_usages,omitempty" gorm:"foreignKey:PageId"`
	HeaderFooterUsages []PageHeaderFooterUsageEntity `json:"header_footer_usages,omitempty" gorm:"foreignKey:PageId"`
}

func (PageEntity) TableName() string {
	return "Site.Pages"
} 