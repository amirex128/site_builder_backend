package user_entity

type RoleEntity struct {
	Id   string `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	Name string `json:"name" gorm:"column:Name" faker:"word"`

	// Relationships
	Users       []UserEntity       `json:"users,omitempty" gorm:"many2many:User.RoleUser;foreignKey:Id;joinForeignKey:RoleId;References:Id;joinReferences:UserId"`
	Permissions []PermissionEntity `json:"permissions,omitempty" gorm:"many2many:User.PermissionRoles;foreignKey:Id;joinForeignKey:RoleId;References:Id;joinReferences:PermissionId"`
	Plans       []PlanEntity       `json:"plans,omitempty" gorm:"many2many:User.RolePlan;foreignKey:Id;joinForeignKey:RoleId;References:Id;joinReferences:PlanId"`
}

func (RoleEntity) TableName() string {
	return "User.Roles"
}
