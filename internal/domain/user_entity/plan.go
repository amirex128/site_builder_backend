package user_entity

type PlanEntity struct {
	Id               string `json:"id" gorm:"column:Id;primaryKey;autoIncrement" faker:"uuid_digit"`
	Name             string `json:"name" gorm:"column:Name" faker:"word"`
	ShowStatus       string `json:"show_status" gorm:"column:ShowStatus" faker:"oneof: visible, hidden"`
	Description      string `json:"description,omitempty" gorm:"column:Description" faker:"paragraph"`
	Price            int64  `json:"price" gorm:"column:Price" faker:"boundary_start=1000, boundary_end=1000000"`
	DiscountType     string `json:"discount_type,omitempty" gorm:"column:DiscountType" faker:"oneof: percent, fixed"`
	Discount         int64  `json:"discount,omitempty" gorm:"column:Discount" faker:"boundary_start=0, boundary_end=100"`
	Duration         int    `json:"duration" gorm:"column:Duration" faker:"boundary_start=1, boundary_end=36"`
	Feature          string `json:"feature,omitempty" gorm:"column:Feature" faker:"sentence"`
	SmsCredits       int    `json:"sms_credits" gorm:"column:SmsCredits" faker:"boundary_start=0, boundary_end=1000"`
	EmailCredits     int    `json:"email_credits" gorm:"column:EmailCredits" faker:"boundary_start=0, boundary_end=1000"`
	StorageMbCredits int    `json:"storage_mb_credits" gorm:"column:StorageMbCredits" faker:"boundary_start=100, boundary_end=10000"`
	AiCredits        int    `json:"ai_credits" gorm:"column:AiCredits" faker:"boundary_start=0, boundary_end=1000"`
	AiImageCredits   int    `json:"ai_image_credits" gorm:"column:AiImageCredits" faker:"boundary_start=0, boundary_end=1000"`
	
	// Relationships
	Roles            []RoleEntity `json:"roles,omitempty" gorm:"many2many:User.RolePlan;foreignKey:Id;joinForeignKey:PlanId;References:Id;joinReferences:RoleId"`
}

func (PlanEntity) TableName() string {
	return "User.Plans"
} 