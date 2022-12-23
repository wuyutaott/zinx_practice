package ziface

type IRouter interface {
	PrevHandle(req IRequest)
	Handle(req IRequest)
	PostHandle(req IRequest)
}
