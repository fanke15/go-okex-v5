package wInterface

import . "github.com/fanke15/go-okex-v5/ws/wImpl"

// 请求数据
type WSParam interface {
	EventType() Event
	ToMap() *map[string]string
}
