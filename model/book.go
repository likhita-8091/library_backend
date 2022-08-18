/*
----------------------------------------------
 File Name: book.go
 Author: jw199@firecloud.ai
 @Company:  Firecloud
 Created Time: 2022/8/18 10:20
-------------------功能说明-------------------
 $
----------------------------------------------
*/

package model

// 图书表
type Book struct {
	CommonModel
	Name       string `json:"name" gorm:"unique"` // 图书名称
	BookKindID uint   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u Book) TableName() string {
	return "book"
}
