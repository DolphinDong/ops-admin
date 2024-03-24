package logic

import (
	"context"
	"github.com/DolphinDong/ops-admin/common/consts"
	"github.com/DolphinDong/ops-admin/common/models/admin"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/DolphinDong/ops-admin/apps/admin/internal/svc"
	pb "github.com/DolphinDong/ops-admin/common/rpc/pb/admin"
)

type GetMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMenuListLogic {
	return &GetMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 列出当前 用户拥有的菜单
func (l *GetMenuListLogic) GetMenuList(req *pb.GetMenuListReq) (*pb.GetMenuListRes, error) {
	db := l.svcCtx.GetDB(l.ctx)
	userId := l.svcCtx.GetUserId(l.ctx)
	logger := l.svcCtx.Logger(l.ctx)
	logger.Info("List user=%v menus", userId)

	result := &pb.GetMenuListRes{}
	var user *admin.User
	var dbMenus []*admin.Menu
	err := db.Model(&admin.User{}).Preload("Roles", "identify = ?", consts.AdminRoleIdentify).Where("status =? and id = ?", admin.UserStatusEnable, userId).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在或者被禁用")
		}
		logger.Errorf("查询用户信息失败：%+v", errors.WithStack(err))
		return nil, errors.WithMessage(err, "查询用户信息失败")
	}
	// 用户为管理员 直接查询所有可用的菜单
	if len(user.Roles) > 0 {
		err = db.Model(&admin.Menu{}).Where("status=? and type in ?", admin.MenuStatusEnable, []int{admin.MenuTypeDir, admin.MenuTypeMenu}).Order("order_no asc").Find(&dbMenus).Error
		if err != nil {
			logger.Errorf("查询用户菜单失败: %+v", errors.WithStack(err))
			return nil, errors.WithMessage(err, "查询用户菜单失败")
		}
	} else { // 不是管理员 查询用户---> 角色 ---> 菜单
		err = db.Model(&admin.User{}).Preload("Roles", "status=?", admin.RoleStatusEnable).
			Preload("Roles.Menus", "status=? and type in ?", admin.MenuStatusEnable, []int{admin.MenuTypeDir, admin.MenuTypeMenu}, func(tx *gorm.DB) *gorm.DB {
				return tx.Order("order_no asc")
			}).
			Where("id=?", userId).Find(&user).Error
		if err != nil {
			logger.Errorf("查询用户菜单失败: %+v", errors.WithStack(err))
			return nil, errors.WithMessage(err, "查询用户菜单失败")
		}
		if user == nil || len(user.Roles) == 0 {
			return result, nil
		}
		dbMenus = l.GetMenus(user)
	}
	menus := l.BuildMenu(dbMenus)
	result.Menu = menus
	return result, nil
}
func (l *GetMenuListLogic) GetMenus(user *admin.User) (menus []*admin.Menu) {
	var hasAppendMenu = make(map[int]struct{})
	for _, role := range user.Roles {
		for _, menu := range role.Menus {
			// 当这个菜单没有添加进menus时需要添加进去，否则无需添加
			if _, exist := hasAppendMenu[menu.Id]; !exist {
				menus = append(menus, menu)
				// 记录menus中已经添加过这个菜单
				hasAppendMenu[menu.Id] = struct{}{}
			}
		}
	}
	return menus
}

func (l *GetMenuListLogic) BuildMenu(dbMenus []*admin.Menu) []*pb.Menu {
	pbMenu := l.ConvertDbMenu2PbMenu(dbMenus)
	return l.FindChildrenMenu(pbMenu, 0)
}

func (l *GetMenuListLogic) FindChildrenMenu(menus []*pb.Menu, parentId int64) []*pb.Menu {
	var result []*pb.Menu
	for _, menu := range menus {
		if menu.ParentMenu == parentId {
			childrenMenu := l.FindChildrenMenu(menus, menu.Id)
			menu.Children = childrenMenu
			result = append(result, menu)
		}
	}
	return result
}

// ConvertDbMenu2PbMenu
//
//	@Description: 将数据模型的menu转换成pb中的menu
//	@receiver l
//	@param dbMenus
//	@return []*pb.Menu
func (l *GetMenuListLogic) ConvertDbMenu2PbMenu(dbMenus []*admin.Menu) []*pb.Menu {
	var result []*pb.Menu
	for _, dbMenu := range dbMenus {
		pbMenu := &pb.Menu{
			Id:         int64(dbMenu.Id),
			Path:       dbMenu.Path,
			Name:       dbMenu.Name,
			Component:  dbMenu.Component,
			Redirect:   dbMenu.Redirect,
			ParentMenu: int64(dbMenu.ParentMenu),
			Meta: &pb.MenuMeta{
				Title:               dbMenu.Title,
				IgnoreKeepAlive:     dbMenu.IgnoreKeepAlive,
				Icon:                dbMenu.Icon,
				HideChildrenInMenu:  dbMenu.HideChildrenInMenu,
				HideMenu:            dbMenu.HideMenu,
				OrderNo:             dbMenu.OrderNo,
				IgnoreRoute:         dbMenu.IgnoreRoute,
				HidePathForChildren: dbMenu.HidePathForChildren,
			},
		}
		result = append(result, pbMenu)
	}
	return result
}
