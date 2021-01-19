package userDao

import (
	"orderbento/src/dao"
	"orderbento/src/models"
	"time"

	"github.com/jinzhu/gorm"
)

type User models.User

func db() *gorm.DB {
	return dao.GetDB()
}

func (u User) TableName() string {
	return "user"
}

func QueryUserByName(name string) (u User) {
	db().Find(&u, "name = ?", name)
	return
}

func QueryUser(pageNo, pageSize int, data map[string]interface{}) (users []User, count int) {
	// ndb := db //不可影響到原本的db (因後面做的 ndb = ndb.where 會改動原本的值)
	ndb := db()
	offset := (pageNo - 1) * pageSize
	if len(data) != 0 {
		ndb = ndb.Where(data)
	}
	ndb.Model(&User{}).Count(&count)
	if count != 0 { //若count為0則不去計算
		ndb.Offset(offset).Limit(pageSize).Find(&users)
	}
	return
}

func (u User) Insert() uint {
	result := db().Create(&u)
	if result.Error != nil {
		panic(result.Error)
	}
	return u.ID
}

func (u User) Update() {
	db().Save(u)
	// fmt.Println(db.Save(u).RowsAffected) // 列出更改的行數
	// db.Model(&User{}).Where("id = ?", u.ID).Update(u) //自定義查詢條件
}

func (u User) Delete() {
	db().Delete(u)
}

/* 登入使用 */
func (u User) UpdateLoginTime() {
	t := time.Now()
	updateData := User{
		LoginTime: &t,
		SessionId: u.SessionId,
	}
	db().Model(&User{}).Where("id = ?", u.ID).Update(updateData)
}
