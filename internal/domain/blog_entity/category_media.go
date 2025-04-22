package blog_entity

type CategoryMediaEntity struct {
	Id         string `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	CategoryId string `json:"category_id" gorm:"column:CategoryId" faker:"uuid_digit"`
	MediaId    string `json:"media_id" gorm:"column:MediaId" faker:"uuid_digit"`
	
	// Relationships
	Category   CategoryEntity `json:"category" gorm:"foreignKey:CategoryId"`
}

func (CategoryMediaEntity) TableName() string {
	return "Blog.CategoryMedia"
} 