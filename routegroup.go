package frouter

import (
	"fmt"
	"net/http"
	"strings"
)

type RouterGroup struct {
	urlPrefix	string
	*http.ServeMux
	middwares	[]Middleware
}

func (g *RouterGroup) GET(path string,handler http.HandlerFunc) {
	g.Request("GET",path,handler)
}

func (g *RouterGroup) POST(path string,handler http.HandlerFunc) {
	g.Request("POST",path,handler)
}

func (g *RouterGroup) PUT(path string,handler http.HandlerFunc) {
	g.Request("PUT",path,handler)

}

func (g *RouterGroup) DELETE(path string,handler http.HandlerFunc) {
	g.Request("DELETE",path,handler)
}

func (g *RouterGroup) Group(path string, handler...Middleware) *RouterGroup {
	path = g.joinPath(g.urlPrefix,path)
	return &RouterGroup{path,g.ServeMux,handler}
}

func (g *RouterGroup) Request(method string,path string,handler http.Handler) {
	path = g.joinPath(g.urlPrefix,path)
	for i := 0; i < len(g.middwares); i++ {
		handler = g.middwares[i](handler)
	}
	g.Handle(path,g.checkMethod(method,handler))
}

func (g *RouterGroup) checkMethod(method string,next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter,r *http.Request) {
		if r.Method != strings.ToUpper(method) {
			w.WriteHeader(404)
			fmt.Fprint(w,"not found")
		}else {
			next.ServeHTTP(w,r)
		}
	}
}

func (g *RouterGroup) joinPath(prefix string,path string) string {
	prefixEle := strings.Split(prefix,"/")
	pathEle := strings.Split(path,"/")
	resEle := make([]string,len(prefixEle)+len(pathEle))
	var i int
	for _,v:=range append(prefixEle,pathEle...){
		v = strings.TrimSpace(v)
		if v != "" {
			resEle[i] = v
			i++
		}
	}
	res := "/" + strings.Join(resEle[:i+1],"/")
	return res
}