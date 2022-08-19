/*
----------------------------------------------
 File Name: book_kind.go
 Author: jw199@firecloud.ai
 @Company:  Firecloud
 Created Time: 2022/8/18 10:06
-------------------功能说明-------------------
 $
----------------------------------------------
*/

package model

import "gorm.io/gorm"

// 图书分类表
type BookKind struct {
	CommonModel
	Name  string `json:"name" gorm:"index;unique"`
	Books []Book `json:"-"`
}

// 用户表，有多张信用卡，类似于图书分类表。
type T1 struct {
	gorm.Model
	CreditCards []CreditCard
}

// 信用卡，类似于图书
type CreditCard struct {
	gorm.Model
	Number string
	T1ID   uint
}

func (u BookKind) TableName() string {
	return "book_kind"
}
