package sdk

// 连接模块
type connect struct {
	serverAddr         string
	sendChan, recvChan chan *Message
}

// 初始化连接模块
func newConnect(serverAddr string) *connect {
	return &connect{
		serverAddr: serverAddr,
		sendChan:   make(chan *Message),
		recvChan:   make(chan *Message),
	}
}

// 发送模块
func (c *connect) send(data *Message) {
	// 直接发送给接收方
	c.recvChan <- data
}

// 接收模块
func (c *connect) recv() <-chan *Message {
	return c.recvChan
}

// 关闭模块
func (c *connect) close() {
	// 目前没啥值得回收的
}
