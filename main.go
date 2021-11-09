package main

import (
	"net/http"

	"github.com/Kevinclane/firstgoproject/controllers"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
