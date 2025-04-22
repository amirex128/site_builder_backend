package blog_entity

type ArticleCategoryEntity struct {
	Id         string `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	ArticleId  string `json:"article_id" gorm:"column:ArticleId" faker:"uuid_digit"`
	CategoryId string `json:"category_id" gorm:"column:CategoryId" faker:"uuid_digit"`
	
	// Relationships
	Article    ArticleEntity  `json:"article" gorm:"foreignKey:ArticleId"`
	Category   CategoryEntity `json:"category" gorm:"foreignKey:CategoryId"`
}

func (ArticleCategoryEntity) TableName() string {
	return "Blog.ArticleCategory"
} 