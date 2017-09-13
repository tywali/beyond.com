package beyond

import (
	"fmt"
	"net/http"
)
type ControllerInterface interface {
	Init(rw http.ResponseWriter, r *http.Request)
	Prepare()
	Get()
	Post()
	Delete()
	Put()
	Head()
	Patch()
	Options()
	Finish()
	/*Render() error
	XSRFToken() string
	CheckXSRFCookie() bool
	HandlerFunc(fn string) bool
	URLMapping()*/
	InitParmsSetting(map[int] string)
}

type Controller struct {
	Rw http.ResponseWriter
	Req *http.Request

	parmNames map[int] string
}

func (c *Controller) Init(rw http.ResponseWriter, r *http.Request) {
	c.Rw = rw
	c.Req = r
}

func (c *Controller) Prepare() {

}
func (c *Controller) Get() {
	fmt.Print("Error get")
}

func (c *Controller) Post() {
	fmt.Print("Error get")
}

func (c *Controller) Delete() {
	fmt.Print("Error get")
}

func (c *Controller) Put() {
	fmt.Print("Error get")
}

func (c *Controller) Head() {
	fmt.Print("Error get")
}

func (c *Controller) Patch() {
	fmt.Print("Error get")
}

func (c *Controller) Options() {
	fmt.Print("Error get")
}

func (c *Controller) Finish() {
	fmt.Print("Error get")
}

func (c *Controller) InitParmsSetting(list map[int] string) {
	c.parmNames = list
}