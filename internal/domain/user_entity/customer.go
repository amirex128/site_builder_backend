package user_entity

import "time"

type CustomerEntity struct {
	Id                 string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	SiteId             string    `json:"site_id" gorm:"column:SiteId" faker:"uuid_digit"`
	FirstName          string    `json:"first_name,omitempty" gorm:"column:FirstName" faker:"first_name"`
	LastName           string    `json:"last_name,omitempty" gorm:"column:LastName" faker:"last_name"`
	AvatarId           string    `json:"avatar_id,omitempty" gorm:"column:AvatarId" faker:"uuid_digit"`
	Email              string    `json:"email" gorm:"column:Email;unique" faker:"email"`
	VerifyEmail        string    `json:"verify_email,omitempty" gorm:"column:VerifyEmail"`
	Password           string    `json:"password" gorm:"column:Password" faker:"-"`
	Salt               string    `json:"salt" gorm:"column:Salt" faker:"-"`
	NationalCode       string    `json:"national_code,omitempty" gorm:"column:NationalCode" faker:"cc_number"`
	Phone              string    `json:"phone,omitempty" gorm:"column:Phone" faker:"phone_number"`
	VerifyPhone        string    `json:"verify_phone,omitempty" gorm:"column:VerifyPhone"`
	IsActive           string    `json:"is_active" gorm:"column:IsActive" faker:"oneof: active, inactive"`
	VerifyCode         int       `json:"verify_code,omitempty" gorm:"column:VerifyCode" faker:"boundary_start=1000, boundary_end=9999"`
	ExpireVerifyCodeAt time.Time `json:"expire_verify_code_at,omitempty" gorm:"column:ExpireVerifyCodeAt" faker:"time"`
	CreatedAt          time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt          time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version            time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted          bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt          time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
	
	// Relationships
	Roles              []RoleEntity   `json:"roles,omitempty" gorm:"many2many:User.CustomerRoles;foreignKey:Id;joinForeignKey:CustomerId;References:Id;joinReferences:RoleId"`
	Addresses          []AddressEntity `json:"addresses,omitempty" gorm:"many2many:User.AddressCustomer;foreignKey:Id;joinForeignKey:CustomerId;References:Id;joinReferences:AddressId"`
}

func (CustomerEntity) TableName() string {
	return "User.Customers"
} 