package handler

import (
	"github.com/gin-gonic/gin"
)

type Todo struct {
	title   string
	context string
}

func NewTodoHandler(title string, context string) (_ *Todo) {
	// 空チェック
	// バリデーション
	return &Todo{title: title, context: context}
}

func AddTodo(c *gin.Context) {
	// request Bodyを取得
	// NewTodoHandlerを呼ぶ
	// DBに要素を追加
	c.JSON(200, "okkkkkk")
}
