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

func (cl *ControllerList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	c := cl.getController(r.URL.Path)

	if c == nil {
		http.NotFound(rw, r)
		return
	}

	c.Init(rw)
	c.Get()
}

func (cl *ControllerList) Add(pattern string, c ControllerInterface) {
	cl.routerMap[pattern] = c
}

func (cl *ControllerList) getController(pattern string) ControllerInterface{
	action := cl.routerMap[pattern]
	return action
}