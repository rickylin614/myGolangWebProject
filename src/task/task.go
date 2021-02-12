package task

import "orderbento/src/task/onlineUserTask"

/* 啟動所有預計的排程 */
func Start() {
	onlineUserTask.StartMemberCheck()
}
