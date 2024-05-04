package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"sync"
)

var (
	Upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	MsgChan = make(chan Message) // 定义全局的消息通道

	// 建立一个 WebSocket 连接的映射，并使用互斥锁进行保护
	Conns = struct {
		sync.RWMutex
		M map[*websocket.Conn]string
	}{M: make(map[*websocket.Conn]string)}
)

type Message struct {
	Sender    string
	Recipient string // 接收人
	Content   string
}
