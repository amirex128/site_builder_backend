package user_entity

import "time"

type CityEntity struct {
	Id         string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	Name       string    `json:"name" gorm:"column:Name" faker:"city"`
	Slug       string    `json:"slug" gorm:"column:Slug" faker:"slug"`
	Status     string    `json:"status" gorm:"column:Status" faker:"oneof: active, inactive"`
	ProvinceId string    `json:"province_id" gorm:"column:ProvinceId" faker:"uuid_digit"`
	Version    time.Time `json:"version" gorm:"column:Version" faker:"time"`

	// Relationships
	Province  ProvinceEntity  `json:"province" gorm:"foreignKey:ProvinceId"`
	Addresses []AddressEntity `json:"addresses,omitempty" gorm:"foreignKey:CityId"`
}

func (CityEntity) TableName() string {
	return "User.Cities"
}
