package main

import (
	"GoWebServer/config"
	"GoWebServer/docs"
	"GoWebServer/server"
	"flag"
	"fmt"
	"os"
)

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = ""

	env := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage : server -e {mode}")
		os.Exit((1))
	}
	flag.Parse()
	config.Init(*env)
	server.Init()
}
