package main

import (
	"fmt"
	"net/http"

	"github.com/sam-karis/basic-web-service/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

func startWebServer(port int, numberOfRetries int) {
	fmt.Println("Starting Web server...")
	fmt.Println("Stoping Web server...")
}
