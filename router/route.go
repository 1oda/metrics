package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"metrics/controller"
	"metrics/utils/response"
	"metrics/utils/route"
)

type router struct {
	engine *gin.Engine
}

type group struct {
	engine      *gin.Engine
	path        string
	middlewares []HandlerFunc
}

type HandlerFunc func(*route.RouteContext)

type method int

const (
	GET    method = 0x000000
	POST   method = 0x000001
	PUT    method = 0x000002
	DELETE method = 0x000003
	ANY    method = 0x000004
)

func newRouter(engine *gin.Engine) *router {
	return &router{
		engine: engine,
	}
}

func (r *router) Group(path string, callback func(group), middlewares ...HandlerFunc) {
	callback(group{
		engine:      r.engine,
		path:        path,
		middlewares: middlewares,
	})

}

func (g group) Group(path string, callback func(group), middlewares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
	g.path += path
	callback(g)
}

func (g group) Registered(method method, url string, action func(ctx *route.RouteContext) *response.Response, middlewares ...HandlerFunc) {

	var handlers = make([]gin.HandlerFunc, len(g.middlewares)+len(middlewares)+1)
	g.middlewares = append(g.middlewares, middlewares...)

	//将中间件转换为gin.HandlerFunc
	for key, middleware := range g.middlewares {
		temp := middleware
		handlers[key] = func(c *gin.Context) {
			temp(&route.RouteContext{Context: c})
		}
	}

	handlers[len(g.middlewares)] = convert(action)
	finalUrl := g.path + url

	switch method {

	case GET:
		g.engine.GET(finalUrl, handlers...)

	case POST:
		g.engine.POST(finalUrl, handlers...)

	case PUT:
		g.engine.PUT(finalUrl, handlers...)

	case DELETE:
		g.engine.DELETE(finalUrl, handlers...)

	case ANY:
		g.engine.Any(finalUrl, handlers...)

	}
}

func Load(r *gin.Engine) {

	router := newRouter(r)
	router.Group("/", func(g group) {
		g.Registered(GET, "/", controller.Index)
	})

	// 监控相关数据源路由
	router.Group("/metrics", func(monitor group) {
		monitor.engine.GET("/metrics", gin.WrapH(promhttp.HandlerFor(controller.Registry(), promhttp.HandlerOpts{})))
	})

	// 进程相关数据源路由
	router.Group("/api", func(process group) {
		process.Registered(GET, "/process", controller.FetchProcess)
		process.Registered(DELETE, "/process", controller.DeleteProcess)
		process.Registered(PUT, "/process", controller.UpdateProcess)
		process.Registered(POST, "/process", controller.AddProcess)

		process.Registered(GET, "/runningProcess", controller.FetchAllRunningProcess)
		process.Registered(GET, "/autoCreateProcess", controller.AutoCreateProcess)
	})
}

func convert(f func(ctx *route.RouteContext) *response.Response) gin.HandlerFunc {
	return func(c *gin.Context) {
		resp := f(&route.RouteContext{
			Context: c,
		})
		data := resp.GetData()
		switch item := data.(type) {
		case string:
			c.String(200, item)
		case gin.H:
			c.JSON(200, item)
		}
	}
}
