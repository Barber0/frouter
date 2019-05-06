package frouter

import (
	"net/http"
)

type IRouter interface {
	Group(path string,handler...http.HandlerFunc) *RouterGroup
	GET(path string,callback http.HandlerFunc)
	POST(path string,callback http.HandlerFunc)
	PUT(path string,callback http.HandlerFunc)
	DELETE(path string,callback http.HandlerFunc)
}

type Router struct {
	RouterGroup
}

func NewFRouter() *Router {
	router := &Router{RouterGroup{
		ServeMux:	http.NewServeMux(),
		urlPrefix:	"/",
	}}
	return router
}