package admin

import "github.com/DolphinDong/ops-admin/common/models"

const (
	RoleStatusEnable  = 1
	RoleStatusDisable = 2
)

type Role struct {
	models.BseId
	Name     string  `gorm:"comment:角色名称;size:255;not null"`
	Identify string  `gorm:"唯一标识;size:50;not null;unique"`
	Status   uint    `gorm:"角色状态;default:1;size:1"`
	Menus    []*Menu `gorm:"many2many:admin_role_menu;constraint:OnDelete:CASCADE"`
	models.ControlBy
	models.ModelTime
}

func (Role) TableName() string {
	return "admin_role"
}
