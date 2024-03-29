// Package controller 聊天控制层
package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"ness_monster/model"

	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
)

// Node userid 和 Node 的映射关系
type Node struct {
	Conn *websocket.Conn
	// 并行转串行,
	DataQueue chan []byte
	GroupSets set.Interface
}

var (
	// 映射关系表
	clientMap = make(map[int64]*Node, 0)
	// 读写锁
	rwlocker sync.RWMutex
)

// Chat 聊天 ws://127.0.0.1/chat?id=1&token=xxxx
func Chat(writer http.ResponseWriter, request *http.Request) {
	// todo 检验接入是否合法
	// checkToken(userId int64,token string)
	// request.URL.Query() 获取url中所有数据
	query := request.URL.Query()
	id := query.Get("id")
	token := query.Get("token")
	// 获取的数据都是字符串，需要做整型
	userID, _ := strconv.ParseInt(id, 10, 64)
	isvalida := checkToken(userID, token)

	// 在webscoket中有处理方法
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// 获得 conn
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}

	comIds := contactService.SearchComunityIds(userID)
	for _, v := range comIds {
		node.GroupSets.Add(v)
	}

	// userid 和 node 形成绑定关系
	rwlocker.Lock() // 操作数据量比较大，所以添加了读写锁
	clientMap[userID] = node
	rwlocker.Unlock()

	// 完成发送逻辑
	go sendproc(node)

	// 完成接收逻辑
	go recvproc(node)

	log.Printf("<-%d\n", userID)
	sendMsg(userID, []byte("hello,world!"))
}

// 发送协程
func sendproc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

// 接收协程
func recvproc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		// 对data进一步处理
		fmt.Printf("recv<=%s", data)
		dispatch(data)
	}
}

// 发送消息
func sendMsg(userID int64, msg []byte) {
	rwlocker.RLock() // 读写锁为了保证并发的安全性
	node, ok := clientMap[userID]
	rwlocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}

// 检测是否有效
func checkToken(userID int64, token string) bool {
	// 从数据库里面查询并比对
	user := userService.Find(userID)
	return user.Token == token
}

// 解析
func dispatch(data []byte) {
	// todo 转成message对象
	// todo 根据cmd参数处理逻辑
	msg := new(model.Message)
	err := json.Unmarshal(data, msg)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	switch msg.Cmd {
	case model.CmdSingleMsg: // 如果是单对单消息,直接将消息转发出去
		sendMsg(msg.Dstid, data)
	case model.CmdRoomMsg: // 群聊消息,需要知道
		// todo 群聊转发逻辑
		for _, v := range clientMap {
			if v.GroupSets.Has(msg.Dstid) {
				v.DataQueue <- data
			}
		}
	case model.CmdHeart: // 心跳事件，保证网络的持久性，如果接到数据说明数据是正常的
		// 啥也别做
	}
}

// AddGroupID 加群
func AddGroupID(userID, gid int64) {
	// 取得 node
	rwlocker.Lock()
	node, ok := clientMap[userID]
	if ok {
		node.GroupSets.Add(gid)
	}
	rwlocker.Unlock()
}
