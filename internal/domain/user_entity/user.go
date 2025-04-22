package user_entity

import (
	"time"
)

type UserEntity struct {
	Id                       string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	FirstName                string    `json:"first_name" gorm:"column:FirstName" faker:"first_name"`
	LastName                 string    `json:"last_name" gorm:"column:LastName" faker:"last_name"`
	Email                    string    `json:"email" gorm:"column:Email;unique" faker:"email"`
	AvatarId                 string    `json:"avatar_id,omitempty" gorm:"column:AvatarId"`
	VerifyEmail              string    `json:"verify_email,omitempty" gorm:"column:VerifyEmail"`
	Password                 string    `json:"password,omitempty" gorm:"column:Password" faker:"-"`
	Salt                     string    `json:"salt,omitempty" gorm:"column:Salt" faker:"-"`
	NationalCode             string    `json:"national_code,omitempty" gorm:"column:NationalCode" faker:"cc_number"`
	Phone                    string    `json:"phone,omitempty" gorm:"column:Phone" faker:"phone_number"`
	VerifyPhone              string    `json:"verify_phone,omitempty" gorm:"column:VerifyPhone"`
	IsActive                 string    `json:"is_active" gorm:"column:IsActive" faker:"oneof: active, inactive"`
	AiTypeEnum               string    `json:"ai_type_enum" gorm:"column:AiTypeEnum" faker:"oneof: free, premium, enterprise"`
	UserTypeEnum             string    `json:"user_type_enum" gorm:"column:UserTypeEnum" faker:"oneof: admin, user, customer"`
	PlanId                   string    `json:"plan_id,omitempty" gorm:"column:PlanId"`
	PlanStartedAt            time.Time `json:"plan_started_at,omitempty" gorm:"column:PlanStartedAt" faker:"time"`
	PlanExpiredAt            time.Time `json:"plan_expired_at,omitempty" gorm:"column:PlanExpiredAt" faker:"time"`
	VerifyCode               int       `json:"verify_code,omitempty" gorm:"column:VerifyCode" faker:"boundary_start=1000, boundary_end=9999"`
	ExpireVerifyCodeAt       time.Time `json:"expire_verify_code_at,omitempty" gorm:"column:ExpireVerifyCodeAt" faker:"time"`
	AiCredits                int       `json:"ai_credits" gorm:"column:AiCredits" faker:"boundary_start=0, boundary_end=1000"`
	AiImageCredits           int       `json:"ai_image_credits" gorm:"column:AiImageCredits" faker:"boundary_start=0, boundary_end=1000"`
	StorageMbCredits         int       `json:"storage_mb_credits" gorm:"column:StorageMbCredits" faker:"boundary_start=100, boundary_end=10000"`
	StorageMbCreditsExpireAt time.Time `json:"storage_mb_credits_expire_at,omitempty" gorm:"column:StorageMbCreditsExpireAt" faker:"time"`
	EmailCredits             int       `json:"email_credits" gorm:"column:EmailCredits" faker:"boundary_start=0, boundary_end=1000"`
	SmsCredits               int       `json:"sms_credits" gorm:"column:SmsCredits" faker:"boundary_start=0, boundary_end=1000"`
	UseCustomEmailSmtp       string    `json:"use_custom_email_smtp" gorm:"column:UseCustomEmailSmtp" faker:"oneof: true, false"`
	SmtpHost                 string    `json:"smtp_host,omitempty" gorm:"column:Smtp_Host" faker:"url"`
	SmtpPort                 int       `json:"smtp_port,omitempty" gorm:"column:Smtp_Port" faker:"boundary_start=0, boundary_end=65535"`
	SmtpUsername             string    `json:"smtp_username,omitempty" gorm:"column:Smtp_Username" faker:"username"`
	SmtpPassword             string    `json:"smtp_password,omitempty" gorm:"column:Smtp_Password" faker:"-"`
	SmtpEnableSsl            bool      `json:"smtp_enable_ssl,omitempty" gorm:"column:Smtp_EnableSsl" faker:"oneof: true, false"`
	SmtpSenderEmail          string    `json:"smtp_sender_email,omitempty" gorm:"column:Smtp_SenderEmail" faker:"email"`
	IsAdmin                  bool      `json:"is_admin" gorm:"column:IsAdmin" faker:"oneof: true, false"`
	CreatedAt                time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt                time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version                  time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted                bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt                time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`

	// Relationships
	Roles     []RoleEntity    `json:"roles,omitempty" gorm:"many2many:User.RoleUser;foreignKey:Id;joinForeignKey:UserId;References:Id;joinReferences:RoleId"`
	Addresses []AddressEntity `json:"addresses,omitempty" gorm:"many2many:User.AddressUser;foreignKey:Id;joinForeignKey:UserId;References:Id;joinReferences:AddressId"`
}

func (UserEntity) TableName() string {
	return "User.Users"
}
