package service

import (
	"fmt"
	"github.com/SuanCaiYv/my-app-backend/db"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"time"
)

type WSApi interface {
	ImageFetch(context *gin.Context)
}

type WSApiHandler struct {
	upgrader  *websocket.Upgrader
	gridFSDao db.GridFSDao
	logger    *logrus.Logger
}

var instanceWSApiHandler *WSApiHandler
var onceWSApiHandler sync.Once

func NewWSApiHandler() *WSApiHandler {
	onceWSApiHandler.Do(newWSApiHandler)
	return instanceWSApiHandler
}

func newWSApiHandler() {
	upgrader := &websocket.Upgrader{}
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	gridFSDao := db.NewGridFSDaoService()
	logger := util.NewLogger()
	instanceWSApiHandler = &WSApiHandler{
		upgrader,
		gridFSDao,
		logger,
	}
}

func (w *WSApiHandler) ImageFetch(context *gin.Context) {
	usernameInterface, ok := context.Get("username")
	var username string
	if !ok {
		username = "test"
	} else {
		username = usernameInterface.(string)
	}
	documentId := context.Query("document_id")
	fmt.Println(documentId)
	request := context.Request
	responseWriter := context.Writer
	socket, err := w.upgrader.Upgrade(responseWriter, request, nil)
	if err != nil {
		w.logger.Errorf("WebSocket专用连接: %v; %s", err, username)
		return
	}
	// 下次无操作关闭时间
	timer := time.After(time.Duration(time.Now().UnixMilli()+int64(2+time.Hour/time.Millisecond)) * time.Millisecond)
	flag := false
	closed := false
	for {
		select {
		case <-timer:
			if flag {
				timer = time.After(time.Duration(time.Now().UnixMilli()+int64(2+time.Hour/time.Millisecond)) * time.Millisecond)
				flag = false
			} else {
				w.logger.Infof("连接闲置过长: %s", username)
				return
			}
		default:
			messageType, bytes, err := socket.ReadMessage()
			if err != nil {
				if !closed {
					w.logger.Errorf("读取失败: %v; %s", err, username)
				} else {
					w.logger.Errorf("连接关闭: %v; %s", err, username)
				}
				return
			}
			switch messageType {
			case websocket.CloseMessage:
				w.logger.Errorf("试图关闭连接: %s", username)
				closed = true
				return
			case websocket.PingMessage:
				w.logger.Infof("ping msg: %s", username)
				err := socket.WriteControl(websocket.PongMessage, make([]byte, 0, 0), time.UnixMilli(int64(15*time.Second/time.Millisecond)+time.Now().UnixMilli()))
				if err != nil {
					w.logger.Errorf("写入Pong失败: %s", username)
					return
				}
			case websocket.PongMessage:
				w.logger.Infof("pong msg: %s", username)
			case websocket.TextMessage:
				// 不应该走到这里
				w.logger.Infof("text msg: %s", username)
				fmt.Println(string(bytes))
			case websocket.BinaryMessage:
				w.logger.Infof("binary msg: %s", username)
			}
		}
	}
}
