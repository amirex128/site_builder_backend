package user_entity

type PermissionEntity struct {
	Id    string `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	Name  string `json:"name" gorm:"column:Name" faker:"word"`
	
	// Relationships
	Roles []RoleEntity `json:"roles,omitempty" gorm:"many2many:User.PermissionRoles;foreignKey:Id;joinForeignKey:PermissionId;References:Id;joinReferences:RoleId"`
}

func (PermissionEntity) TableName() string {
	return "User.Permissions"
} 