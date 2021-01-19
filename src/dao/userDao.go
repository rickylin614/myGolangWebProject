package dao

import (
	"orderbento/src/models"
	"time"
)

type User models.User

func (u User) TableName() string {
	return "user"
}

func QueryUserByName(name string) (u User) {
	db.Find(&u, "name = ?", name)
	return
}

func QueryUser(data map[string]interface{}) (users []User, count int) {
	ndb := db //不可影響到原本的db (因後面做的 ndb = ndb.where 會改動原本的值)
	pageNo := 1
	pageSize := 20
	if val, ok := data["pageNo"].(int); ok {
		pageNo = val
	}
	if val, ok := data["pageSize"].(int); ok {
		pageSize = val
	}
	if val, ok := data["name"].(string); ok && val != "" {
		ndb = ndb.Where("name = ?", val)
	}
	offset := (pageNo - 1) * pageSize
	db.Count(&count)
	if count != 0 { //若count為0則不去計算
		ndb.Offset(offset).Limit(pageSize).Find(&users)
	}
	return
}

func (u User) Insert() uint {
	result := db.Create(&u)
	if result.Error != nil {
		panic(result.Error)
	}
	return u.ID
}

func (u User) Update() {
	db.Save(u)
	// fmt.Println(db.Save(u).RowsAffected) // 列出更改的行數
	// db.Model(&User{}).Where("id = ?", u.ID).Update(u) //自定義查詢條件
}

func (u User) Delete() {
	db.Delete(u)
}

/* 登入使用 */
func (u User) UpdateLoginTime() {
	t := time.Now()
	updateData := User{
		LoginTime: &t,
		SessionId: u.SessionId,
	}
	db.Model(&User{}).Where("id = ?", u.ID).Update(updateData)
}
