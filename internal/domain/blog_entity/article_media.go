package blog_entity

type ArticleMediaEntity struct {
	Id        string `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	ArticleId string `json:"article_id" gorm:"column:ArticleId" faker:"uuid_digit"`
	MediaId   string `json:"media_id" gorm:"column:MediaId" faker:"uuid_digit"`
	
	// Relationships
	Article   ArticleEntity `json:"article" gorm:"foreignKey:ArticleId"`
}

func (ArticleMediaEntity) TableName() string {
	return "Blog.ArticleMedia"
} 