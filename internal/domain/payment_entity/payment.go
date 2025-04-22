package payment_entity

import "time"

type PaymentEntity struct {
	Id                  string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	SiteId              string    `json:"site_id" gorm:"column:SiteId" faker:"uuid_digit"`
	PaymentStatusEnum   string    `json:"payment_status_enum" gorm:"column:PaymentStatusEnum" faker:"oneof: pending, successful, failed"`
	UserType            string    `json:"user_type,omitempty" gorm:"column:UserType" faker:"oneof: normal, vip"`
	TrackingNumber      int64     `json:"tracking_number" gorm:"column:TrackingNumber" faker:"boundary_start=1000000, boundary_end=9999999"`
	Gateway             string    `json:"gateway" gorm:"column:Gateway" faker:"oneof: zarinpal, mellat, saman, parsian, pasargad"`
	GatewayAccountName  string    `json:"gateway_account_name" gorm:"column:GatewayAccountName" faker:"word"`
	Amount              int64     `json:"amount" gorm:"column:Amount" faker:"boundary_start=1000, boundary_end=1000000"`
	ServiceName         string    `json:"service_name" gorm:"column:ServiceName" faker:"word"`
	ServiceAction       string    `json:"service_action" gorm:"column:ServiceAction" faker:"word"`
	OrderId             int64     `json:"order_id" gorm:"column:OrderId" faker:"boundary_start=1000, boundary_end=9999"`
	ReturnUrl           string    `json:"return_url" gorm:"column:ReturnUrl" faker:"url"`
	CallVerifyUrl       string    `json:"call_verify_url" gorm:"column:CallVerifyUrl" faker:"url"`
	ClientIp            string    `json:"client_ip" gorm:"column:ClientIp" faker:"ipv4"`
	Message             string    `json:"message,omitempty" gorm:"column:Message" faker:"sentence"`
	GatewayResponseCode string    `json:"gateway_response_code,omitempty" gorm:"column:GatewayResponseCode" faker:"oneof: 0, 1, -1"`
	TransactionCode     string    `json:"transaction_code,omitempty" gorm:"column:TransactionCode" faker:"uuid_digit"`
	AdditionalData      string    `json:"additional_data,omitempty" gorm:"column:AdditionalData" faker:"paragraph"`
	OrderData           string    `json:"order_data,omitempty" gorm:"column:OrderData" faker:"paragraph"`
	UserId              string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CustomerId          string    `json:"customer_id" gorm:"column:CustomerId" faker:"uuid_digit"`
	CreatedAt           time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt           time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version             time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted           bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt           time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
}

func (PaymentEntity) TableName() string {
	return "Payment.Payments"
} 