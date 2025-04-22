package site_entity

type PageHeaderFooterUsageEntity struct {
	Id             string `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	PageId         string `json:"page_id" gorm:"column:PageId" faker:"uuid_digit"`
	HeaderFooterId string `json:"header_footer_id" gorm:"column:HeaderFooterId" faker:"uuid_digit"`
	SiteId         string `json:"site_id" gorm:"column:SiteId" faker:"uuid_digit"`
	UserId         string `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	
	// Relationships
	Page          PageEntity          `json:"page" gorm:"foreignKey:PageId"`
	HeaderFooter  HeaderFooterEntity `json:"header_footer" gorm:"foreignKey:HeaderFooterId"`
	Site          SiteEntity         `json:"site" gorm:"foreignKey:SiteId"`
}

func (PageHeaderFooterUsageEntity) TableName() string {
	return "Site.PageHeaderFooterUsages"
} 