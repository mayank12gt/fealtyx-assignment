package main

import (
	"flag"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/mayank12gt/fealtyx_assignment/repository"
	"github.com/mayank12gt/fealtyx_assignment/service"
)

type Config struct {
	Port            string
	Ollama_Base_Url string
}

type App struct {
	StudentService *service.StudentService
	Config         Config
}

func main() {

	var url string
	var port string

	//get configuration from cmd flags
	flag.StringVar(&url, "Ollama_URL", "", "Ollama Base Url")
	flag.StringVar(&port, "Port", "4000", "Port no. for API")
	flag.Parse()

	if url == "" {
		log.Println("Warning: Ollama_URL is empty")
	}

	// Initialize the repository with dummy data
	studentRepo := &repository.StudentRepo{
		Students: []repository.Student{
			{
				ID:    1,
				Name:  "Mayank Gupta",
				Age:   21,
				Email: "mayank.gt15@gmail.com",
			},
			{
				ID:    2,
				Name:  "John Doe",
				Age:   24,
				Email: "johndoe@gmail.com",
			},
			{
				ID:    3,
				Name:  "Jane Doe",
				Age:   25,
				Email: "janedoe@gmail.com",
			},
		},
	}

	// Initialize the service with the repository
	studentService := &service.StudentService{
		StudentRepo: studentRepo,
	}

	//create App object to encapsulate dependencies and configs and implement dependency injection
	app := App{
		StudentService: studentService,
		Config: Config{
			Port:            port,
			Ollama_Base_Url: url,
		},
	}

	app.serve()

}

// create the server
func (app *App) serve() {

	server := echo.New()

	app.registerHandlers(server)

	server.Start(":" + app.Config.Port)

}

// register handlers
func (app *App) registerHandlers(server *echo.Echo) {

	server.GET("/", app.HealthCheckHandler())

	server.GET("/students", app.GetStudentsHandler())

	server.POST("/students", app.CreateStudentHandler())

	server.GET("/students/:id", app.GetStudentHandler())

	server.DELETE("/students/:id", app.DeleteStudentHandler())

	server.PUT("/students/:id", app.UpdateStudentHandler())

	server.GET("/students/:id/summary", app.GetStudentSummaryHandler())

}
