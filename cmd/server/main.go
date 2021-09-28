package main

import "fmt"

//Application configurations
type App struct {
}

//Sets up the application.
func (app *App) Run() error {
	fmt.Println("Setting up Applicaiton")

	return nil
}

func main() {
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up Rest API")
		fmt.Println(err)
	}
}
