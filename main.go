package main

import (
	"fmt"
	"net/http"

	"github.com/shridharkalagi/dummy/controllers"
)

func main() {
	fmt.Println("From a module")

	name := "shridhar"
	fmt.Println(name)

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
