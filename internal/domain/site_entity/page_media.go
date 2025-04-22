package site_entity

type PageMediaEntity struct {
	Id      string `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	PageId  string `json:"page_id" gorm:"column:PageId" faker:"uuid_digit"`
	MediaId string `json:"media_id" gorm:"column:MediaId" faker:"uuid_digit"`
	
	// Relationships
	Page    PageEntity `json:"page" gorm:"foreignKey:PageId"`
}

func (PageMediaEntity) TableName() string {
	return "Site.PageMedia"
} 