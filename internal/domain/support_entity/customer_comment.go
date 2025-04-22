package support_entity

import "time"

type CustomerCommentEntity struct {
	Id               string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	CustomerTicketId string    `json:"customer_ticket_id" gorm:"column:CustomerTicketId" faker:"uuid_digit"`
	Content          string    `json:"content" gorm:"column:Content" faker:"paragraph"`
	RespondentId     string    `json:"respondent_id" gorm:"column:RespondentId" faker:"uuid_digit"`
	CreatedAt        time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version          time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted        bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt        time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
	
	// Relationships
	CustomerTicket   CustomerTicketEntity `json:"customer_ticket" gorm:"foreignKey:CustomerTicketId"`
}

func (CustomerCommentEntity) TableName() string {
	return "Support.CustomerComments"
} 