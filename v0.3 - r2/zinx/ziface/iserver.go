package ziface

type IServer interface {
	Start()
	Stop()
	Serve()
	SetRouter(router IRoute)
}