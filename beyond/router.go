package beyond

import (
	"net/http"
	"strings"
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
	ctx := new(Context)
	ctx.Init()
	idx := 0
	path := "/"
	semi := strings.Split(pattern, "/")
	for _, v := range semi {
		if strings.HasPrefix(v, ":") {
			ctx.AddParam(idx, v[1:len(v)])
			idx++
		} else {
			path += v
		}
	}
	c.SetContext(ctx)
	cl.routerMap[path] = c
}

func (cl *ControllerList) getController(pattern string) ControllerInterface{
	idx := -1
	path := "/"
	semi := strings.Split(pattern, "/")
	found := false
	var action ControllerInterface
	for _, v := range semi {
		path += v
		if !found {
			action = cl.routerMap[path]
			if action != nil {
				found = true
			}
		}
		if found {
			action.GetContext().AddParamVal(idx, v)
			idx++
		}
	}
	return action
}
