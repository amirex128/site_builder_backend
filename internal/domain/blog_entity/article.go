package blog_entity

import (
	"site_builder_backend/internal/domain/site_entity"
	"time"
)

type ArticleEntity struct {
	Id           string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	Title        string    `json:"title,omitempty" gorm:"column:Title" faker:"sentence"`
	Description  string    `json:"description,omitempty" gorm:"column:Description" faker:"paragraph"`
	Body         string    `json:"body,omitempty" gorm:"column:Body" faker:"paragraph"`
	Slug         string    `json:"slug" gorm:"column:Slug" faker:"slug"`
	SiteId       string    `json:"site_id" gorm:"column:SiteId" faker:"uuid_digit"`
	VisitedCount int       `json:"visited_count" gorm:"column:VisitedCount" faker:"boundary_start=0, boundary_end=10000"`
	ReviewCount  int       `json:"review_count" gorm:"column:ReviewCount" faker:"boundary_start=0, boundary_end=500"`
	Rate         int       `json:"rate" gorm:"column:Rate" faker:"boundary_start=0, boundary_end=5"`
	Badges       string    `json:"badges,omitempty" gorm:"column:Badges" faker:"sentence"`
	SeoTags      string    `json:"seo_tags,omitempty" gorm:"column:SeoTags" faker:"sentence"`
	UserId       string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version      time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted    bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt    time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`

	// Relationships
	Categories []CategoryEntity                     `json:"categories,omitempty" gorm:"many2many:Blog.ArticleCategory;foreignKey:Id;joinForeignKey:ArticleId;References:Id;joinReferences:CategoryId"`
	Media      []ArticleMediaEntity                 `json:"media,omitempty" gorm:"foreignKey:ArticleId"`
	PageUsages []site_entity.PageArticleUsageEntity `json:"page_usages,omitempty" gorm:"foreignKey:ArticleId"`
}

func (ArticleEntity) TableName() string {
	return "Blog.Articles"
}
