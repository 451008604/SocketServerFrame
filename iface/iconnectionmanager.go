package iface

// 连接管理器
type IConnectionManager interface {
	// 生成一个新的连接Id
	NewConnId() int
	// 遍历所有连接
	RangeConnections(handler func(conn IConnection))
	// 添加连接
	Add(conn IConnection)
	// 删除连接
	Remove(conn IConnection)
	// 根据ConnId获取连接
	Get(connId int) (IConnection, bool)
	// 获取当前连接数量
	Len() int
	// 删除并停止所有连接
	ClearConn()

	// 设置连接创建时的Hook函数
	SetConnOnOpened(connOpenCallBack func(connection IConnection))
	// 执行连接创建时的Hook函数
	ConnOnOpened(connection IConnection)

	// 设置连接断开时的Hook函数
	SetConnOnClosed(connCloseCallBack func(connection IConnection))
	// 执行连接断开时的Hook函数
	ConnOnClosed(connection IConnection)

	// 触发限流时的Hook函数
	SetConnOnRateLimiting(limitCallBack func(connection IConnection))
	// 执行触发限流时的Hook函数
	ConnRateLimiting(connection IConnection)
}
