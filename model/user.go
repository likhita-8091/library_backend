package model

import (
	"golang.org/x/crypto/bcrypt"
)

type Role int

const (
	Admin       Role = iota + 1 // 超级管理员
	BookManager                 // 图书管理员
	Reader                      // 读者
)

type User struct {
	CommonModel
	Name       string `json:"name,omitempty" gorm:"unique" validate:"required"`             // 用户名
	NickName   string `json:"nick_name"`                                                    // 昵称
	Role       Role   `json:"role,omitempty"gorm:"<-:create" validate:"required,gt=1,lt=4"` // 角色
	Sex        int    `json:"sex,omitempty"`                                                // 性别
	Classed    string `json:"classed,omitempty"`                                            // 班级
	Code       string `json:"code,omitempty"`                                               // 学号
	Phone      string `json:"phone,omitempty"`                                              // 手机号
	Password   string `json:"-"`                                                            // 密码
	FirstLogin bool   `gorm:"default:true" json:"first_login,omitempty"`                    // 第一次登陆
}

func (u User) TableName() string {
	return "user_info"
}

// 创建一个admin用户
func NewAdmin() *User {
	admin := &User{
		Name:       "admin",
		NickName:   "江不凡",
		Password:   "admin",
		Role:       Admin,
		FirstLogin: true,
	}
	// 密码加密
	admin.Password = EncryptPassword(admin.Password)
	return admin
}

// BcryptHash 使用 bcrypt 对密码进行加密
func EncryptPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

// DecryptPassword 对比明文密码和数据库的哈希值
func DecryptPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
