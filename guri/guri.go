package guri

import (
	"log"
	"net/http"
)

// 提供给框架用户，用来定义路由映射的处理方法
type HandlerFunc func(*Context)

type (
	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup
	}
	RouterGroup struct {
		prefix string
		engine *Engine
		parent *RouterGroup
		middlewares []HandlerFunc
	}
)

// 框架的构造函数
func New() *Engine {
	engine := &Engine {router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

// 创建新的子RouterGroup
// 所有的group都引用同一个Engine实例
func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup {
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup)Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares ...)
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}