package support_entity

import "time"

type TicketEntity struct {
	Id         string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	Title      string    `json:"title" gorm:"column:Title" faker:"sentence"`
	Status     string    `json:"status" gorm:"column:Status" faker:"oneof: open, pending, closed"`
	Category   string    `json:"category" gorm:"column:Category" faker:"oneof: technical, billing, general"`
	AssignedTo string    `json:"assigned_to,omitempty" gorm:"column:AssignedTo" faker:"uuid_digit"`
	ClosedBy   string    `json:"closed_by,omitempty" gorm:"column:ClosedBy" faker:"uuid_digit"`
	ClosedAt   time.Time `json:"closed_at,omitempty" gorm:"column:ClosedAt" faker:"time"`
	Priority   string    `json:"priority" gorm:"column:Priority" faker:"oneof: low, medium, high, urgent"`
	UserId     string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version    time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted  bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt  time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
	
	// Relationships
	Comments   []CommentEntity   `json:"comments,omitempty" gorm:"foreignKey:TicketId"`
	Media      []TicketMediaEntity `json:"media,omitempty" gorm:"foreignKey:TicketId"`
}

func (TicketEntity) TableName() string {
	return "Support.Tickets"
} 