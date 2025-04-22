package support_entity

type CustomerTicketMediaEntity struct {
	Id               string `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	CustomerTicketId string `json:"customer_ticket_id" gorm:"column:CustomerTicketId" faker:"uuid_digit"`
	MediaId          string `json:"media_id" gorm:"column:MediaId" faker:"uuid_digit"`
	
	// Relationships
	CustomerTicket   CustomerTicketEntity `json:"customer_ticket" gorm:"foreignKey:CustomerTicketId"`
}

func (CustomerTicketMediaEntity) TableName() string {
	return "Support.CustomerTicketMedia"
} 