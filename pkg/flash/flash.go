// Package flash 用以支持在会话中存储消息提示
package flash

import (
	"encoding/gob"
	"goblog/pkg/session"
)

// Flashes Flash 消息数组类型，用以在会话中存储 map
type Flashes map[string]interface{}

// 存入会话数据里的 key

var flashKey = "_flashes"

func init() {
	// 在 gorilla/sessions 中存储 map 和 struct 数据需
	// 要提前注册 gob，方便后续 gob 序列化编码、解码
	//标准库 gob 是 Go 专属的编解码方式，是标准库自带的一个数据结构序列化的编码 / 解码工具。类似于 JSON 或 XML，不过执行效率比他们更高。特别适合在 Go 语言程序间传递数据。

	gob.Register(Flashes{})
}

// 私有方法，新增一条提示
func addFlash(key string, message string) {
	flashkes := Flashes{}
	flashkes[key] = message
	session.Put(flashKey, flashkes)
	session.Save()
}

// Info 添加 Info 类型的消息提示
func Info(message string) {
	addFlash("info", message)
}

// Warning 添加 Warning 类型的消息提示
func Warning(message string) {
	addFlash("warning", message)
}

// Success 添加 Success 类型的消息提示
func Success(message string) {
	addFlash("success", message)
}

// Danger 添加 Danger 类型的消息提示
func Danger(message string) {
	addFlash("danger", message)
}

// All 获取所有消息
// All 获取所有消息
func All() Flashes {
	val := session.Get(flashKey)
	// 读取时必须做类型检测
	flashMessages, ok := val.(Flashes)
	if !ok {
		return nil
	}
	// 读取即销毁，直接删除
	session.Forget(flashKey)
	return flashMessages
}
