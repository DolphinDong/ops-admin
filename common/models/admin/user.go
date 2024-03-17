package admin

import (
	"github.com/DolphinDong/ops-admin/common/models"
	"gorm.io/gorm"
)

const (
	UserStatusEnable  = 1
	UserStatusDisable = 2
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"comment:用户登录名;not null;unique;size:50"`
	NickName string `json:"nickName" gorm:"comment:用户名;not null;size:50"`
	Password string `json:"password" gorm:"comment:密码;not null"`
	Email    string `json:"email" gorm:"comment:邮箱;size:50"`
	Tel      string `json:"tel" gorm:"comment:手机号;size:11"`
	Status   int    `json:"status" gorm:"comment:账号状态，1启用，2禁言;default:1;size:1"`
	models.ControlBy
}

func (User) TableName() string {
	return "admin_user"
}
