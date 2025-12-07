package sdk

// MsgTypeText 相关常量信息
const (
	MsgTypeText = "text"
)

// Chat 聊天结构体
type Chat struct {
	Nick      string
	UserID    string
	SessionID string
	conn      *connect
}

// Message 消息结构体
type Message struct {
	Type       string
	Name       string
	FormUserID string
	ToUserID   string
	Content    string
	Session    string
}

// NewChat 聊天结构体构造函数
func NewChat(serverAddr, nick, userID, sessionID string) *Chat {
	return &Chat{
		Nick:      nick,
		UserID:    userID,
		SessionID: sessionID,
		conn:      newConnect(serverAddr),
	}
}

// Send 发送消息
func (chat *Chat) Send(msg *Message) {
	chat.conn.send(msg)
}

// Close 关闭聊天
func (chat *Chat) Close() {
	chat.conn.close()
}

// Recv 接受消息
func (chat *Chat) Recv() <-chan *Message {
	return chat.conn.recv()
}
