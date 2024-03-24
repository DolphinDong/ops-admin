package admin

import "github.com/DolphinDong/ops-admin/common/models"

const (
	MenuTypeDir    = 0
	MenuTypeMenu   = 1
	MenuTypeButton = 2
)
const (
	MenuStatusEnable  = 0
	MenuStatusDisable = 1
)

type Menu struct {
	models.BseId
	Path                string  `gorm:"comment:路由路径"`
	Name                string  `gorm:"comment:路由名称;unique;not null;size:255"`
	Component           string  `gorm:"comment:菜单组件"`
	Redirect            string  `gorm:"comment:重定向路径"`
	Title               string  `gorm:"comment:菜单标题;not null;size:255"`
	IgnoreKeepAlive     bool    `gorm:"comment:是否忽略KeepAlive缓存"`
	Icon                string  `gorm:"comment:图标;size:100"`
	HideChildrenInMenu  bool    `gorm:"comment:隐藏所有子菜单"`
	HideMenu            bool    `gorm:"comment:前路由不在菜单显示"`
	OrderNo             int64   `gorm:"comment:菜单排序，只对第一级有效"`
	IgnoreRoute         bool    `gorm:"comment:忽略路由"`
	HidePathForChildren bool    `gorm:"comment:是否在子级菜单的完整path中忽略本级path"`
	Type                int64   `gorm:"comment:菜单类型"`
	Permission          string  `gorm:"comment:权限标识;size:100"`
	Status              int     `gorm:"comment:菜单状态"`
	ParentMenu          uint    `gorm:"comment:父级菜单"`
	Children            []*Menu `gorm:"-"`
	Roles               []*Role `gorm:"many2many:admin_role_menu;constraint:OnDelete:CASCADE"`
	Apis                []*Api  `gorm:"many2many:admin_menu_api;constraint:OnDelete:CASCADE"`
	models.ModelTime
	models.ControlBy
}

func (Menu) TableName() string {
	return "admin_menu"
}
