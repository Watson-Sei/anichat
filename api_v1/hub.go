package main

import (
	"encoding/json"
	"github.com/Watson-Sei/anichat/api_v1/plugin"
	"github.com/gofiber/websocket/v2"
	"log"
)

type message struct {
	conn *websocket.Conn
	data []byte
	room string
}

type subscription struct {
	conn *websocket.Conn
	room string
}

type hub struct {
	rooms map[string]map[*websocket.Conn]bool
	broadcast chan message
	register chan subscription
	unregister chan subscription
}

var h = hub{
	rooms: make(map[string]map[*websocket.Conn]bool),
	broadcast: make(chan message),
	register: make(chan subscription),
	unregister: make(chan subscription),
}

// room -> connection -> token, name etc...
var Member = make(map[string]map[*websocket.Conn]map[string]interface{})

var MemberCount int = 1

func (h *hub) run() {
	for {
		select {
		case s := <- h.register:
			connections := h.rooms[s.room]
			if connections == nil {
				connections = make(map[*websocket.Conn]bool)
				h.rooms[s.room] = connections
			}
			h.rooms[s.room][s.conn] = true
			log.Println("connection registered")

			// create token
			token := plugin.MakeToken(MemberCount, SecretKey)

			// add Member
			Member[s.room] = make(map[*websocket.Conn]map[string]interface{})
			Member[s.room][s.conn] = make(map[string]interface{})
			Member[s.room][s.conn]["token"] = token
			Member[s.room][s.conn]["name"] = nil
			Member[s.room][s.conn]["count"] = MemberCount

			// Member count ++
			MemberCount += 1

			bytes, err := json.Marshal(map[string]interface{}{
				"event": "token",
				"token": token,
			})
			if err != nil {
				log.Println(err)
			}

			if s.conn.WriteMessage(websocket.TextMessage, bytes); err != nil {
				log.Println("write error:", err)

				delete(connections, s.conn)
				if len(connections) == 0 {
					delete(h.rooms, s.room)
				}
			}

		case s := <- h.unregister:
			connections := h.rooms[s.room]
			if connections != nil {
				if _, ok := connections[s.conn]; ok {
					delete(connections, s.conn)
					log.Println("connection unregistered")
				}
			}

		case m := <- h.broadcast:
			log.Println("message received:", string(m.data))
			connections := h.rooms[m.room]

			var dataMap map[string]interface{}
			if err := json.Unmarshal(m.data, &dataMap); err != nil {
				log.Println("error unmarshal: ", err)
			}

			if dataMap["event"] == "post" {

				bytes, err := json.Marshal(map[string]interface{}{
					"event": "member-post",
					"token": dataMap["token"],
					"name": dataMap["name"],
					"message": dataMap["message"],
				})
				if err != nil {
					log.Println(err)
				}

				for connection := range connections {
					if err = connection.WriteMessage(websocket.TextMessage, bytes); err != nil {
						log.Println("write error:", err)

						delete(connections, connection)
						if len(connections) == 0 {
							delete(h.rooms, m.room)
						}
					}
				}

				log.Println("member-postされました")
			}

			if dataMap["event"] == "join" {
				// トークンが正しければ
				if authToken(m.conn, m.room, dataMap["token"].(string)) {
					// 入室OK + 現在の入室一覧を通知
					memberlist := getMemberList(m.room)

					bytes, err := json.Marshal(map[string]interface{}{
						"event": "join-result",
						"status": true,
						"list": memberlist,
					})
					if err != nil {
						log.Println(err)
					}

					if err = m.conn.WriteMessage(websocket.TextMessage, bytes); err != nil {
						log.Println("write error: ", err)

						delete(connections, m.conn)
						if len(connections) == 0 {
							delete(h.rooms, m.room)
						}
					}

					// メンバー一覧に追加
					Member[m.room][m.conn]["name"] = dataMap["name"]

					// 入室通知
					dataMap["event"] = "member-join"
					// 本人へ送信
					bytes, err = json.Marshal(dataMap)
					if err != nil {
						log.Println(err)
					}
					if m.conn.WriteMessage(websocket.TextMessage, bytes); err != nil {
						log.Println("write error: ", err)

						delete(connections, m.conn)
						if len(connections) == 0 {
							delete(h.rooms, m.room)
						}
					}
					// 本人以外へ送信
					dataMap["token"] = Member[m.room][m.conn]["count"]
					bytes, err = json.Marshal(dataMap)
					if err != nil {
						log.Println(err)
					}
					for connection := range connections {
						log.Println("順番接続:", connection)
						if m.conn != connection {
							log.Println("実行接続:", connection)
							if connection.WriteMessage(websocket.TextMessage, bytes); err != nil {
								log.Println(err)

								delete(connections, connection)
								if len(connections) == 0 {
									delete(h.rooms, m.room)
								}
							}
						}
					}
				} else {
					// トークンが誤っていた場合
					// 本人にNG通知
					bytes, err := json.Marshal(map[string]interface{}{
						"event": "join-result",
						"status": false,
					})
					if err != nil {
						log.Println(err)
					}

					if err = m.conn.WriteMessage(websocket.TextMessage, bytes); err != nil {
						log.Println("write error: ", err)

						delete(connections, m.conn)
						if len(connections) == 0 {
							delete(h.rooms, m.room)
						}
					}
				}
			}
		}
	}
}

func authToken(conn *websocket.Conn, room,token string) bool {
	if _, ok := Member[room][conn]; ok {
		if token == Member[room][conn]["token"] {
			return true
		}
		return false
	}
	return false
}

func getMemberList(room string) []map[string]interface{} {
	var list []map[string]interface{}
	for key, _ := range Member[room] {
		cur := Member[room][key]
		if cur["name"] != nil {
			list = append(list, map[string]interface{}{
				"token": cur["count"],
				"name": cur["name"],
			})
		}
	}
	return list
}