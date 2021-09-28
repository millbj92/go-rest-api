package main

import (
	"fmt"
	"net/http"

	transportHttp "github.com/millbj92/go-rest-api/internal/transport/http"
)

//Application configurations
type App struct {
}

//Sets up the application.
func (app *App) Run() error {
	fmt.Println("Setting up Applicaiton")

	handler := transportHttp.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failedto set up server")
		return err
	}
	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up Rest API")
		fmt.Println(err)
	}
}
