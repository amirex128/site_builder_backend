package user_entity

type ProvinceEntity struct {
	Id     string `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	Name   string `json:"name" gorm:"column:Name" faker:"state"`
	Slug   string `json:"slug" gorm:"column:Slug" faker:"slug"`
	Status int    `json:"status" gorm:"column:Status" faker:"oneof: 1, 2, 3"`

	// Relationships
	Cities    []CityEntity    `json:"cities,omitempty" gorm:"foreignKey:ProvinceId"`
	Addresses []AddressEntity `json:"addresses,omitempty" gorm:"foreignKey:ProvinceId"`
}

func (ProvinceEntity) TableName() string {
	return "User.Provinces"
}
