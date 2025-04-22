package ai_entity

import "time"

type CreditEntity struct {
	Id         string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	UserId     string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CustomerId string    `json:"customer_id" gorm:"column:CustomerId" faker:"uuid_digit"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version    time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted  bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt  time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
}

func (CreditEntity) TableName() string {
	return "Ai.Credits"
}
