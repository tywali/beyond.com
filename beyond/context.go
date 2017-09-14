package beyond

type Context struct {
	parmSettings map[int] string
	Parms map[string] string
}

func (c *Context) Init() {
	c.parmSettings = make(map[int] string)
	c.Parms = make(map[string] string)
}

func (c *Context) AddParam(idx int, name string)  {
	c.parmSettings[idx] = name
}

func (c *Context) AddParamVal(idx int, val string) {
	name := c.parmSettings[idx]
	c.Parms[name] = val
}