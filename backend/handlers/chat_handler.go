package handlers

import (
	"fmt"
	"github.com/gorilla/websocket"
	. "main/backend/models"
	"net/http"
	"strings"
)

//func Chat() {
//	for {
//		msg := <-MsgChan // 改为 Message 类型
//
//		// 将消息存储到数据库中
//		insertQuery := "INSERT INTO messages (sender, recipient, message) VALUES (?, ?, ?)"
//		_, err := db.Exec(insertQuery, msg.Sender, msg.Recipient, msg.Content)
//		if err != nil {
//			fmt.Printf("Failed to insert message to database: %v", err)
//			continue
//		}
//
//		// 根据接收人发送消息
//		Conns.RLock()
//		for conn, username := range Conns.M {
//			if username == msg.Recipient {
//				if err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Content)); err != nil {
//					fmt.Printf("Failed to send message: %+v\n", err)
//				}
//				break
//			}
//		}
//		Conns.RUnlock()
//	}
//}

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

	// 在连接建立后发送在线用户列表
	sendUserList()

	defer func() {
		// 当连接断开时，从映射中移除
		Conns.Lock()
		delete(Conns.M, conn)
		Conns.Unlock()

		// 在连接断开后发送在线用户列表
		sendUserList()
	}()

	// 读取消息信息
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}

		// 解析接收人和消息内容
		//parts := strings.SplitN(string(msg), ":", 1)
		//if len(parts) != 2 {
		//	fmt.Printf("%s\n", string(msg))
		//	continue
		//}
		recipient := "zjy"
		content := string(msg)
		fmt.Printf("%s\n", content)
		message := Message{
			Sender:    username,
			Recipient: recipient,
			Content:   fmt.Sprintf("%s (%s): %s", username, ip, content),
		}

		// 将消息发送到全局消息通道
		MsgChan <- message
	}
}

func sendUserList() {
	Conns.RLock()
	defer Conns.RUnlock()

	var userList []string
	for _, username := range Conns.M {
		userList = append(userList, username)
	}

	userListMessage := strings.Join(userList, ", ")
	for conn := range Conns.M {
		if err := conn.WriteMessage(websocket.TextMessage, []byte("Online users: "+userListMessage)); err != nil {
			fmt.Printf("Failed to send user list: %+v\n", err)
		}
	}
}
