package site_entity

import "time"

type SiteEntity struct {
	Id         string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	Domain     string    `json:"domain" gorm:"column:Domain" faker:"domain_name"`
	DomainType string    `json:"domain_type" gorm:"column:DomainType" faker:"oneof: custom, subdomain"`
	Name       string    `json:"name" gorm:"column:Name" faker:"company"`
	Status     string    `json:"status" gorm:"column:Status" faker:"oneof: active, inactive, pending"`
	SiteType   string    `json:"site_type" gorm:"column:SiteType" faker:"oneof: blog, ecommerce, portfolio, business"`
	UserId     string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version    time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted  bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt  time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
	
	// Relationships
	Settings   *SettingsEntity `json:"settings,omitempty" gorm:"foreignKey:SiteId"`
	Pages      []PageEntity    `json:"pages,omitempty" gorm:"foreignKey:SiteId"`
}

func (SiteEntity) TableName() string {
	return "Site.Sites"
} 