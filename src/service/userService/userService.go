package userService

import (
	"orderbento/src/dao/userDao"
	"orderbento/src/utils"
)

/* 查詢用戶名稱 */
func QueryUserByName(name string) (u userDao.User) {
	return userDao.QueryUserByName(name)
}

/* 條件式查詢用戶 */
func QueryUser(data map[string]interface{}) (users []userDao.User, count int) {
	pageNo, pageSize := utils.GetPage(data)
	params := make(map[string]interface{})
	if val, ok := data["name"].(string); ok && val != "" {
		params["name"] = val
	}
	return userDao.QueryUser(pageNo, pageSize, params)
}

/* 新增 */
func Insert(user userDao.User) uint {
	return user.Insert()
}

/* 修改 */
func Update(user userDao.User) {
	user.Update()
}

/* 刪除 */
func Delete(user userDao.User) {
	user.Delete()
}

/* 登入使用 */
func UpdateLoginTime(user userDao.User) {
	user.UpdateLoginTime()
}
