package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UID      uint   `json:"uid,omitempty" gorm:"not null;unique"`
	Name     string `json:"name" gorm:"type:varchar(20);not null; comment:'用户名'"`
	Password string `json:"password" gorm:"type:varchar(255);not null; comment:'密码'"`
	Mobile   string `json:"mobile" gorm:"type:varchar(11);not null;unique;comment:'手机号'"`
	Email    string `json:"email" gorm:"type:varchar(255);not null;unique;comment:'邮箱'"`
	RoleId   int    `json:"role_id" gorm:"type:int;not null; comment:'角色';default:0"`
	Status   int    `json:"status" gorm:"type:TINYINT(1);not null; comment:'状态：1启用 0禁用';default:1"`
}
