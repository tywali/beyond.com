package beyond

import (
	"net/http"
)

type ControllerList struct {
	routerMap map[string] ControllerInterface
}

func (cl *ControllerList) Init() {
	cl.routerMap = make(map[string] ControllerInterface)
}

func execController(c ControllerInterface, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		c.Get()
	case http.MethodPost:
		c.Post()
	case http.MethodDelete:
		c.Delete()
	case http.MethodPut:
		c.Put()
	case http.MethodHead:
		c.Head()
	case http.MethodPatch:
		c.Patch()
	case http.MethodOptions:
		c.Options()
	}
}

func (cl *ControllerList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	c := cl.getController(r.URL.Path)

	if c == nil {
		http.NotFound(rw, r)
		return
	}

	c.Init(rw, r)
	execController(c, r)
}

func (cl *ControllerList) Add(pattern string, c ControllerInterface) {
	cl.routerMap[pattern] = c
}

func (cl *ControllerList) getController(pattern string) ControllerInterface{
	action := cl.routerMap[pattern]
	return action
}