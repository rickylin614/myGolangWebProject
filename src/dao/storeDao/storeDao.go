package storeDao

import (
	"orderbento/src/dao"
	"orderbento/src/models"

	"github.com/jinzhu/gorm"
)

type Store models.Store

func db() *gorm.DB {
	return dao.GetDB()
}

func (s Store) TableName() string {
	return "store"
}

func QueryStoreByName(name string) (s Store) {
	db().Find(&s, "name = ?", name)
	return
}

func QueryStoreById(id uint) (s Store) {
	db().Find(&s, "id = ?", id)
	return
}

func QueryStore(pageNo, pageSize int, data map[string]interface{}) (stores []Store, count int) {
	// ndb := db //不可影響到原本的db (因後面做的 ndb = ndb.where 會改動原本的值)
	ndb := db()
	offset := (pageNo - 1) * pageSize
	if len(data) != 0 {
		ndb = ndb.Where(data)
	}
	ndb.Model(&Store{}).Count(&count)
	if count != 0 { //若count為0則不去計算
		ndb.Offset(offset).Limit(pageSize).Find(&stores)
	}
	return
}

func (s Store) Insert() uint {
	result := db().Create(&s)
	if result.Error != nil {
		panic(result.Error)
	}
	return s.ID
}

func (s Store) Update() {
	db().Save(s)
	// fmt.Println(db.Save(u).RowsAffected) // 列出更改的行數
	// db.Model(&User{}).Where("id = ?", u.ID).Update(u) //自定義查詢條件
}

func (s Store) Delete() {
	db().Delete(s)
}
