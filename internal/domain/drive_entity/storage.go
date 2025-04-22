package drive_entity

import "time"

type StorageEntity struct {
	Id          string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	UsedSpaceKb int64     `json:"used_space_kb" gorm:"column:UsedSpaceKb" faker:"boundary_start=1000, boundary_end=1000000"`
	QuotaKb     int64     `json:"quota_kb" gorm:"column:QuotaKb" faker:"boundary_start=1000000, boundary_end=10000000"`
	ChargedAt   time.Time `json:"charged_at" gorm:"column:ChargedAt" faker:"time"`
	ExpireAt    time.Time `json:"expire_at" gorm:"column:ExpireAt" faker:"time"`
	UserId      string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version     time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted   bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt   time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`
}

func (StorageEntity) TableName() string {
	return "Drive.Storages"
}
 