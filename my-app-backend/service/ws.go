package service

import (
	"fmt"
	"github.com/SuanCaiYv/my-app-backend/db"
	"github.com/SuanCaiYv/my-app-backend/util"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
	"sync"
	"time"
)

type WSApi interface {
	ImageFetch(context *gin.Context)

	Generic(context *gin.Context)
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
	username := context.MustGet("username").(string)
	articleId := context.Query("article_id")
	fmt.Println(articleId)
	request := context.Request
	responseWriter := context.Writer
	socket, err := w.upgrader.Upgrade(responseWriter, request, nil)
	if err != nil {
		w.logger.Errorf("WebSocket专用连接: %v; %s", err, username)
		return
	}
	socket.SetPingHandler(func(appData string) error {
		err := socket.WriteControl(websocket.PongMessage, []byte(appData), time.Now().Add(5*time.Second))
		if err == websocket.ErrCloseSent {
			return nil
		} else if e, ok := err.(net.Error); ok && e.Temporary() {
			return nil
		}
		return err
	})
	socket.SetPongHandler(func(appData string) error {
		return nil
	})
	socket.SetCloseHandler(func(code int, text string) error {
		w.logger.Infof("关闭连接: %d, %s", code, text)
		err := socket.WriteControl(websocket.CloseMessage, []byte("close"), time.Now().Add(5*time.Second))
		if err != nil {
			w.logger.Error(err)
			return err
		}
		return nil
	})
	lastTimestamp := time.Now()
	isActive := false
	// 闲置时间
	finish := time.After(time.Duration(time.Now().Add(2*time.Hour).UnixMilli() * int64(time.Millisecond)))
	for {
		select {
		case <-finish:
			if !isActive {
				return
			} else {
				isActive = false
				finish = time.After(time.Duration(lastTimestamp.Add(2*time.Hour).UnixMilli() * int64(time.Millisecond)))
			}
		default:
			lastTimestamp = time.Now()
			isActive = true
			messageType, bytes, err := socket.ReadMessage()
			if err != nil {
				w.logger.Error(err)
				return
			}
			switch messageType {
			case websocket.PingMessage:
				_ = socket.PingHandler()("pong")
			case websocket.PongMessage:
				_ = socket.PongHandler()("")
			case websocket.CloseMessage:
				_ = socket.CloseHandler()(0, "")
			case websocket.TextMessage:
				// 处理文本消息
				w.logger.Infof("%s: %s", username, string(bytes))
			case websocket.BinaryMessage:
				// 处理二进制消息
				w.logger.Infof("%s: %d", username, len(bytes))
			}
		}
	}
}

func (w *WSApiHandler) Generic(context *gin.Context) {
	username := context.MustGet("username").(string)
	socket, err := w.upgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		w.logger.Error(err)
		return
	}
	socket.SetPingHandler(func(appData string) error {
		err := socket.WriteControl(websocket.PongMessage, []byte(appData), time.Now().Add(5*time.Second))
		if err == websocket.ErrCloseSent {
			return nil
		} else if e, ok := err.(net.Error); ok && e.Temporary() {
			return nil
		}
		return err
	})
	socket.SetPongHandler(func(appData string) error {
		return nil
	})
	socket.SetCloseHandler(func(code int, text string) error {
		w.logger.Infof("关闭连接: %d, %s", code, text)
		err := socket.WriteControl(websocket.CloseMessage, []byte("close"), time.Now().Add(5*time.Second))
		if err != nil {
			w.logger.Error(err)
			return err
		}
		return nil
	})
	lastTimestamp := time.Now()
	isActive := false
	// 闲置时间
	finish := time.After(time.Duration(time.Now().Add(2*time.Hour).UnixMilli() * int64(time.Millisecond)))
	for {
		select {
		case <-finish:
			if !isActive {
				return
			} else {
				isActive = false
				finish = time.After(time.Duration(lastTimestamp.Add(2*time.Hour).UnixMilli() * int64(time.Millisecond)))
			}
		default:
			lastTimestamp = time.Now()
			isActive = true
			messageType, bytes, err := socket.ReadMessage()
			if err != nil {
				w.logger.Error(err)
				return
			}
			switch messageType {
			case websocket.PingMessage:
				_ = socket.PingHandler()("pong")
			case websocket.PongMessage:
				_ = socket.PongHandler()("")
			case websocket.CloseMessage:
				_ = socket.CloseHandler()(0, "")
			case websocket.TextMessage:
				// 处理文本消息
				w.logger.Infof("%s: %s", username, string(bytes))
			case websocket.BinaryMessage:
				// 处理二进制消息
				w.logger.Infof("%s: %d", username, len(bytes))
			}
		}
	}
}
