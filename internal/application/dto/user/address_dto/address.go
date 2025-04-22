package address_dto

import "site_builder_backend/internal/domain/user/address_entity"

type CreateAddressDto struct {
}
type UpdateAddressDto struct {
	Id string `json:"id"`
}

func (CreateAddressDto) ToAddressEntity() *address_entity.AddressEntity {
	return nil
}

func (UpdateAddressDto) ToAddressEntity() *address_entity.AddressEntity {
	return nil
}
