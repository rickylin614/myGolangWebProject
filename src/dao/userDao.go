package dao

import (
	"orderbento/src/models"
	"time"
)

// type User struct {
// 	Name      string
// 	Pwd       string
// 	SessionId string
// 	LoginTime time.Time `gorm:"type:time"`
// 	gorm.Model
// }

type User models.User

func (u User) TableName() string {
	return "user"
}

func QueryUserByName(name string) (u User) {
	db.Find(&u, "name = ?", name)
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
