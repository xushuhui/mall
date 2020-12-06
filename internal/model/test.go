package model

import (
	"fmt"
	"mall_go/global"
)

type SysBaseMenu struct {
	ID uint `gorm:"primaryKey"`
	//CreatedAt       time.Time
	//UpdatedAt       time.Time
	//DeletedAt       *time.Time

	MenuLevel uint   `json:"-"`
	ParentId  string `json:"parentId" gorm:"comment:父菜单ID"`
	Path      string `json:"path" gorm:"comment:路由path"`
	Name      string `json:"name" gorm:"comment:路由name"`
	Hidden    bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`
	Component string `json:"component" gorm:"comment:对应前端文件路径"`
	Sort      int    `json:"sort" gorm:"comment:排序标记"`
	//SysAuthoritys []SysAuthority         `json:"authoritys" gorm:"many2many:sys_authority_menu;"`
	//Children      []SysBaseMenu          `json:"children" gorm:"-"`
}
type SysAuthority struct {
	//CreatedAt       time.Time
	//UpdatedAt       time.Time
	//DeletedAt       *time.Time     `sql:"index"`
	ID uint `gorm:"primaryKey"`
	//AuthorityId     string         `json:"authorityId" gorm:"not null;unique;primary_key;comment:角色ID;size:90"`
	AuthorityName string `json:"authorityName" gorm:"comment:角色名"`
	ParentId      string `json:"parentId" gorm:"comment:父角色ID"`
	//DataAuthorityId []SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	//Children        []SysAuthority `json:"children" gorm:"-"`
	Menus []SysBaseMenu `json:"menus" gorm:"many2many:sys_authority_menu;"`
}

func GetA() (s SysAuthority) {
	//var s SysAuthority
	e := global.DBEngine.Preload("SysBaseMenus").First(&s, "id = ?", 888)
	fmt.Println("e", e)
	return
}
