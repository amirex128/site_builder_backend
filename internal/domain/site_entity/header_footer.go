package site_entity

import "time"

type HeaderFooterEntity struct {
	Id        string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	SiteId    string    `json:"site_id" gorm:"column:SiteId" faker:"uuid_digit"`
	Title     string    `json:"title" gorm:"column:Title" faker:"sentence"`
	IsMain    bool      `json:"is_main" gorm:"column:IsMain" faker:"oneof: true, false"`
	Body      string    `json:"body,omitempty" gorm:"column:Body" faker:"paragraph"`
	Type      int       `json:"type" gorm:"column:Type" faker:"oneof: 1, 2"`
	UserId    string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CreatedAt time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version   time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
	
	// Relationships
	Site      SiteEntity `json:"site" gorm:"foreignKey:SiteId"`
	HeaderPages []PageEntity `json:"header_pages,omitempty" gorm:"foreignKey:HeaderId"`
	FooterPages []PageEntity `json:"footer_pages,omitempty" gorm:"foreignKey:FooterId"`
	PageUsages []PageHeaderFooterUsageEntity `json:"page_usages,omitempty" gorm:"foreignKey:HeaderFooterId"`
}

func (HeaderFooterEntity) TableName() string {
	return "Site.HeaderFooters"
} 