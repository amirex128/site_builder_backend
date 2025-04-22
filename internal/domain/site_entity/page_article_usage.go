package site_entity

type PageArticleUsageEntity struct {
	Id        string `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	PageId    string `json:"page_id" gorm:"column:PageId" faker:"uuid_digit"`
	ArticleId string `json:"article_id" gorm:"column:ArticleId" faker:"uuid_digit"`
	SiteId    string `json:"site_id" gorm:"column:SiteId" faker:"uuid_digit"`
	UserId    string `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	
	// Relationships
	Page     PageEntity `json:"page" gorm:"foreignKey:PageId"`
	Site     SiteEntity `json:"site" gorm:"foreignKey:SiteId"`
}

func (PageArticleUsageEntity) TableName() string {
	return "Site.PageArticleUsages"
} 