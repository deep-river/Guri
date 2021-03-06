package guri

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
)

// 自定义的路由映射处理方法
type HandlerFunc func(*Context)

type (
	Engine struct {
		*RouterGroup
		router *router
		groups []*RouterGroup // 路由分组
		htmlTemplates *template.Template
		funcMap template.FuncMap // 自定义template渲染函数
	}
	// 路由组
	RouterGroup struct {
		prefix string
		engine *Engine
		parent *RouterGroup
		middlewares []HandlerFunc // 应用的中间件
	}
)
// 框架的构造函数
func New() *Engine {
	engine := &Engine {router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}
// 默认使用Logger Recovery的框架构造函数
func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
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
// 注册使用中间件
func (group *RouterGroup)Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares ...)
}
// 注册GET路由
func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}
// 注册POST路由
func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}
// 静态文件路由
func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*fielpath")
	group.GET(urlPattern, handler)
}

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath,  http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}
// 设置自定义渲染函数
func (engine *Engine) SetFuncMap(funcMap template.FuncMap) {
	engine.funcMap = funcMap
}
// 加载HTML模板
func (engine *Engine) LoadHTMLGlob(pattern string) {
	engine.htmlTemplates = template.Must(template.New("").Funcs(engine.funcMap).ParseGlob((pattern)))
}

func (engine *Engine) Run(addr string) (err error) {
	// http.ListenAndServe(addr string, handler http.Handler) error
	return http.ListenAndServe(addr, engine)
}

// 实现 Handler接口的 ServeHTTP方法
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares ...)
		}
	}
	c := newContext(w, req)
	c.handlers = middlewares
	c.engine = engine
	engine.router.handle(c)
}

