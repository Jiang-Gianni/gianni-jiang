package server

type ContextKey string

const (
	CtxKeyTodo   ContextKey = "todo"
	CtxKeyStatus ContextKey = "status"
)
