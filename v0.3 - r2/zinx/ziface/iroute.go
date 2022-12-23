package ziface

type IRoute interface {
	PrevHandle(req IRequest)
	Handle(req IRequest)
	PostHandle(req IRequest)
}