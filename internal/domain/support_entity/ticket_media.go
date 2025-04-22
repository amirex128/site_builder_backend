package support_entity

type TicketMediaEntity struct {
	Id       string `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	TicketId string `json:"ticket_id" gorm:"column:TicketId" faker:"uuid_digit"`
	MediaId  string `json:"media_id" gorm:"column:MediaId" faker:"uuid_digit"`
	
	// Relationships
	Ticket   TicketEntity `json:"ticket" gorm:"foreignKey:TicketId"`
}

func (TicketMediaEntity) TableName() string {
	return "Support.TicketMedia"
} 