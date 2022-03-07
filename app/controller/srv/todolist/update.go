package todolist

import (
	"net/http"
	"todolist/app/model/mongo/todolist"

	"github.com/gin-gonic/gin"
)

// Update - 完成待辦事項
func Update(context *gin.Context) {

	var status string
	var msg string

	defer func() {
		context.JSON(http.StatusOK, gin.H{
			"status":  status,
			"message": msg,
		})
	}()

	// 取得資料
	id := context.Param("id")
	memberId := context.GetString("memberId")

	// 確認資料是否存在
	todoList, err := todolist.GetById(id)

	if err != nil {
		status = "failed"
		msg = "更新失敗，待辦事項不存在"
		return
	}

	// 確認使用者是否正確
	if memberId != todoList.MemberId {
		status = "failed"
		msg = "更新失敗，使用者錯誤"
		return
	}

	// 更新資料庫資料
	if err := todolist.Update(id); err != nil {
		status = "failed"
		msg = "更新失敗，資料庫錯誤"
		return
	}

	status = "ok"
	msg = "已成功完成待辦事項"

	return
}
