package ziface

type IServer interface {
	// 启动服务器
	Start()

	// 停止服务器
	Stop()

	// 提供服务
	Serve()
}
