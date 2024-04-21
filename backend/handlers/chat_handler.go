package handlers

import (
	"fmt"
	"github.com/gorilla/websocket"
	. "main/backend/models"
	"net/http"
)

func chat() {
	for {
		msg := <-MsgChan // 改为 Message 类型

		// 将消息存储到数据库中
		insertQuery := "INSERT INTO messages (sender, message) VALUES (?, ?)"
		_, err := db.Exec(insertQuery, msg.Sender, msg.Content)
		if err != nil {
			fmt.Printf("Failed to insert message to database: %v", err)
			continue
		}

		// 广播消息给所有连接的 WebSocket 客户端
		Conns.RLock()
		for conn := range Conns.M {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Content)); err != nil {
				fmt.Printf("Failed to broadcast message: %+v\n", err)
			}
		}
		Conns.RUnlock()
	}
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	// 升级连接为 WebSocket
	conn, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
		return
	}

	// 在这个位置获取客户端IP地址和用户名
	ip := r.RemoteAddr
	cookie, err := r.Cookie("username")
	if err != nil {
		fmt.Println("Error getting username from cookie:", err)
		return
	}
	username := cookie.Value

	// 将新的 WebSocket 连接添加到映射中
	Conns.Lock()
	Conns.M[conn] = username
	Conns.Unlock()

	defer func() {
		// 当连接断开时，从映射中移除
		Conns.Lock()
		delete(Conns.M, conn)
		Conns.Unlock()
	}()

	// 读取消息信息
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		message := Message{
			Sender:  username,
			Content: fmt.Sprintf("%s (%s): %s", username, ip, string(msg)),
		}

		// 将消息发送到全局消息通道
		MsgChan <- message
	}

}
