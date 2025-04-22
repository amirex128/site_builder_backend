package drive_entity

import "time"

type FileItemEntity struct {
	Id          string    `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	Name        string    `json:"name" gorm:"column:Name" faker:"file_name"`
	BucketName  string    `json:"bucket_name" gorm:"column:BucketName" faker:"word"`
	ServerKey   string    `json:"server_key" gorm:"column:ServerKey" faker:"uuid_digit"`
	FilePath    string    `json:"file_path" gorm:"column:FilePath" faker:"file_path"`
	IsDirectory bool      `json:"is_directory" gorm:"column:IsDirectory" faker:"oneof: true, false"`
	Size        int64     `json:"size" gorm:"column:Size" faker:"boundary_start=1000, boundary_end=1000000"`
	MimeType    string    `json:"mime_type" gorm:"column:MimeType" faker:"mime_type"`
	ParentId    string    `json:"parent_id,omitempty" gorm:"column:ParentId" faker:"uuid_digit"`
	Permission  string    `json:"permission" gorm:"column:Permission" faker:"oneof: private, public, shared"`
	UserId      string    `json:"user_id" gorm:"column:UserId" faker:"uuid_digit"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:CreatedAt" faker:"time"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:UpdatedAt" faker:"time"`
	Version     time.Time `json:"version" gorm:"column:Version" faker:"time"`
	IsDeleted   bool      `json:"is_deleted" gorm:"column:IsDeleted" faker:"oneof: true, false"`
	DeletedAt   time.Time `json:"deleted_at,omitempty" gorm:"column:DeletedAt" faker:"time"`

	// Relationships
	Parent   *FileItemEntity  `json:"parent,omitempty" gorm:"foreignKey:ParentId"`
	Children []FileItemEntity `json:"children,omitempty" gorm:"foreignKey:ParentId"`
}

func (FileItemEntity) TableName() string {
	return "Drive.FileItems"
}
 