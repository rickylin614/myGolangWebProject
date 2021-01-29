package storeService

import (
	"orderbento/src/dao/storeDao"
)

/* 查詢用戶名稱 */
func QueryStoreByName(name string) (s storeDao.Store) {
	return storeDao.QueryStoreByName(name)
}

func QueryStoreById(id uint) (s storeDao.Store) {
	return storeDao.QueryStoreById(id)
}

/* 條件式查詢用戶 */
func QueryStore(data map[string]interface{}) (stores []storeDao.Store, count int) {
	pageNo := 1
	pageSize := 20
	if val, ok := data["pageNo"].(float64); ok {
		pageNo = int(val)
	}
	if val, ok := data["pageSize"].(float64); ok {
		pageSize = int(val)
	}
	params := make(map[string]interface{})
	if val, ok := data["name"].(string); ok && val != "" {
		params["name"] = val
	}

	return storeDao.QueryStore(pageNo, pageSize, params)
}

/* 新增 */
func Insert(store storeDao.Store) uint {
	return store.Insert()
}

/* 修改 */
func Update(store storeDao.Store) {
	store.Update()
}

/* 刪除 */
func Delete(store storeDao.Store) {
	store.Delete()
}
