package network

import (
	"fmt"
	"github.com/451008604/socketServerFrame/iface"
	"github.com/451008604/socketServerFrame/logs"
	pb "github.com/451008604/socketServerFrame/proto/bin"
	"sync"
)

type Connection struct {
	Server       iface.IServer          // 当前Conn所属的Server
	ConnID       int                    // 当前连接的ID（SessionID）
	isClosed     bool                   // 当前连接是否已关闭
	MsgHandler   iface.IMsgHandler      // 消息管理MsgId和对应处理函数的消息管理模块
	ExitBuffChan chan bool              // 通知该连接已经退出的channel
	msgBuffChan  chan []byte            // 用于读、写两个goroutine之间的消息通信
	property     map[string]interface{} // 连接属性
	propertyLock sync.RWMutex           // 连接属性读写锁
}

// 启动接收消息协程
func (c *Connection) StartReader() {
}

// 启动发送消息协程
func (c *Connection) StartWriter() {
}

// 启动连接
func (c *Connection) Start() {
	// 将新建的连接添加到所属Server的连接管理器内
	c.Server.GetConnMgr().Add(c)

	// 在收到退出消息时释放进程
	for range c.ExitBuffChan {
		return
	}
}

// 停止连接
func (c *Connection) Stop() {
	if c.isClosed {
		return
	}
	c.isClosed = true
	// 通知关闭该连接的监听
	c.ExitBuffChan <- true

	// 将连接从连接管理器中删除
	c.Server.GetConnMgr().Remove(c)

	// 关闭该连接管道
	close(c.ExitBuffChan)
	close(c.msgBuffChan)
}

// 获取当前连接ID
func (c *Connection) GetConnID() int {
	return c.ConnID
}

func (c *Connection) RemoteAddrStr() string {
	return ""
}

// 发送消息给客户端
func (c *Connection) SendMsg(msgId pb.MessageID, data []byte) {
	if c.isClosed {
		logs.PrintLogInfo(fmt.Sprintf("连接已关闭导致消息发送失败 -> msgId:%v\tdata:%v", msgId, string(data)))
		return
	}

	// 将消息数据封包
	msg := c.Server.DataPacket().Pack(NewMsgPackage(msgId, data))
	if msg == nil {
		return
	}
	// 写入传输通道发送给客户端
	c.msgBuffChan <- msg
}

// 设置连接属性
func (c *Connection) SetProperty(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[key] = value
}

// 获取连接属性
func (c *Connection) GetProperty(key string) interface{} {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	if value, ok := c.property[key]; ok {
		return value
	} else {
		return nil
	}
}

// 删除连接属性
func (c *Connection) RemoveProperty(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}
