package user_entity

import "time"

type AddressEntity struct {
	Id          string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	Title       string    `json:"title,omitempty" gorm:"column:Title" faker:"sentence"`
	Latitude    float32   `json:"latitude,omitempty" gorm:"column:Latitude" faker:"lat"`
	Longitude   float32   `json:"longitude,omitempty" gorm:"column:Longitude" faker:"long"`
	AddressLine string    `json:"address_line" gorm:"column:AddressLine" faker:"address"`
	PostalCode  string    `json:"postal_code" gorm:"column:PostalCode" faker:"zip"`
	CityId      string    `json:"city_id" gorm:"column:CityId" faker:"uuid_digit"`
	ProvinceId  string    `json:"province_id" gorm:"column:ProvinceId" faker:"uuid_digit"`
	UserId      string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CustomerId  string    `json:"customer_id" gorm:"column:CustomerId" faker:"uuid_digit"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version     time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted   bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt   time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`

	// Relationships
	City      CityEntity       `json:"city,omitempty" gorm:"foreignKey:CityId"`
	Province  ProvinceEntity   `json:"province,omitempty" gorm:"foreignKey:ProvinceId"`
	Users     []UserEntity     `json:"users,omitempty" gorm:"many2many:User.AddressUser;foreignKey:Id;joinForeignKey:AddressId;References:Id;joinReferences:UserId"`
	Customers []CustomerEntity `json:"customers,omitempty" gorm:"many2many:User.AddressCustomer;foreignKey:Id;joinForeignKey:AddressId;References:Id;joinReferences:CustomerId"`
}

func (AddressEntity) TableName() string {
	return "User.Addresses"
}
