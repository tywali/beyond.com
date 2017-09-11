package beyond

import (
	"net/http"
	"log"
)

var controllList  = new(ControllerList)

func Router(pattern string, c ControllerInterface) {
	controllList.Add(pattern, c)
}

func Run() {
	err := http.ListenAndServe(":9090", controllList)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}