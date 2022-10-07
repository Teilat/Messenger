package resolver

import (
	"Messenger/database"
	"Messenger/webapi/converters"
	"Messenger/webapi/models"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub
	// The websocket connection.
	conn *websocket.Conn
	// Buffered channel of outbound messages.
	send     chan *WsMessage
	resolver *Resolver
	// user and chat ids for websocket connection
	userId      string
	localChatId string
	chatId      uuid.UUID
}

type WsMessage struct {
	Message []byte
	// chat id for message distribution
	chatId uuid.UUID
}

type Hub struct {
	// Registered clients.
	clients map[*Client]bool
	// Inbound messages from the clients.
	broadcast chan *WsMessage
	// Register requests from the clients.
	register chan *Client
	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan *WsMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, wsErr := c.conn.ReadMessage()
		if wsErr != nil {
			if websocket.IsUnexpectedCloseError(wsErr, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				c.resolver.Log.Printf("error: %v", wsErr)
			}
			break
		}

		var dat models.MessageType
		err := json.Unmarshal(message, &dat)
		if err != nil {
			c.resolver.Log.Printf("error: %v", err)
		} else {
			message = []byte{}
		}

		if dat.Payload == nil {
			c.resolver.Log.Printf("userId=%s, localChatId=%s, Payload empty, message=%s, wsErr=%s", c.userId, c.localChatId, message, wsErr.Error())
			break
		}

		switch dat.Action {
		case "sendMessage":
			msg, err := c.resolver.CreateWsMessage(models.SendMessage{Text: dat.Payload["text"].(string)}, c.localChatId, c.userId)
			if err != nil {
				c.resolver.Log.Printf("error: %v", err)
				break
			}
			//c.resolver.Log.Printf("new message %s", dat.Payload["text"].(string))
			res, err := json.Marshal(converters.MessagesToWsMessages([]database.Message{*msg}))
			message = res
		case "editMessage":
			err := c.resolver.EditWsMessage(models.EditMessage{NewText: dat.Payload["text"].(string), MessageId: uint32(dat.Payload["messageId"].(float64))}, c.localChatId, c.userId)
			if err != nil {
				c.resolver.Log.Printf("error: %v", err)
				break
			}
		case "deleteMessage":
			err := c.resolver.DeleteWsMessage(models.DeleteMessage{MessageId: uint32(dat.Payload["messageId"].(float64))}, c.localChatId, c.userId)
			if err != nil {
				c.resolver.Log.Printf("error: %v", err)
				break
			}
		case "replyMessage":
			msg, err := c.resolver.ReplyWsMessage(models.ReplyMessage{ReplyMessageId: uint32(dat.Payload["replyMessageId"].(float64)), Text: dat.Payload["text"].(string)}, c.localChatId, c.userId)
			if err != nil {
				c.resolver.Log.Printf("error: %v", err)
				break
			}
			//c.resolver.Log.Printf("new message %s", dat.Payload["text"].(string))
			res, err := json.Marshal(converters.MessagesToWsMessages([]database.Message{*msg}))
			message = res
		case "getMessages":
			payload := models.GetMessages{Limit: int(dat.Payload["limit"].(float64)), Offset: int(dat.Payload["offset"].(float64))}
			messages, err := c.resolver.GetWsMessages(payload, c.localChatId, c.userId)
			if err != nil {
				c.resolver.Log.Printf("error: %v", err)
				break
			}
			err = c.conn.WriteJSON(converters.MessagesToWsMessages(messages))
			if err != nil {
				c.resolver.Log.Printf("error: %v", err)
				break
			}
		default:
		}

		if len(message) > 0 {
			c.hub.broadcast <- &WsMessage{Message: message, chatId: c.chatId}
		}
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			if message.chatId == c.chatId {
				w, err := c.conn.NextWriter(websocket.TextMessage)
				if err != nil {
					return
				}
				w.Write(message.Message)

				// Add queued chat messages to the current websocket message.
				n := len(c.send)
				for i := 0; i < n; i++ {
					w.Write(newline)
					msg := <-c.send
					w.Write(msg.Message)
				}

				if err := w.Close(); err != nil {
					return
				}
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// ServeWs handles websocket requests from the peer.
func (r Resolver) ServeWs(hub *Hub, conn *websocket.Conn, res *Resolver, userId string, chatId string) {
	client := &Client{hub: hub, conn: conn, send: make(chan *WsMessage, 256), userId: userId, localChatId: chatId, chatId: res.ChatIdToUUID(chatId, userId), resolver: res}
	client.hub.register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}
