package beyond

import (
	"fmt"
	"net/http"
)
type ControllerInterface interface {
	Init(rw http.ResponseWriter)
	Get()
}

type Controller struct {
	Rw http.ResponseWriter
	rw http.ResponseWriter
}

func (c *Controller) Get() {
	//http.Error(c.Ctx.ResponseWriter, "Method Not Allowed", 405)
	fmt.Print("Error get")
}

func (c *Controller) Init(rw http.ResponseWriter) {
	c.Rw = rw
}
