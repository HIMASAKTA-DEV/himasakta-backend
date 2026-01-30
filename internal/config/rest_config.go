package config

import (
	"fmt"

	"log"
	"os"

	"github.com/Flexoo-Academy/Golang-Template/db"
	"github.com/Flexoo-Academy/Golang-Template/internal/api/controller"
	"github.com/Flexoo-Academy/Golang-Template/internal/api/repository"
	"github.com/Flexoo-Academy/Golang-Template/internal/api/routes"
	"github.com/Flexoo-Academy/Golang-Template/internal/api/service"
	"github.com/Flexoo-Academy/Golang-Template/internal/middleware"
	mailer "github.com/Flexoo-Academy/Golang-Template/internal/pkg/email"
	"github.com/gin-gonic/gin"
)

type RestConfig struct {
	server *gin.Engine
}

func NewRest() RestConfig {
	db := db.New()
	app := gin.Default()
	server := NewRouter(app)
	middleware := middleware.New(db)

	var (
		//=========== (PACKAGE) ===========//
		mailerService mailer.Mailer = mailer.New()
		// oauthService  oauth.Oauth   = oauth.New()
		// awsS3Service  storage.AwsS3 = storage.NewAwsS3()

		//=========== (REPOSITORY) ===========//
		userRepository         repository.UserRepository         = repository.NewUser(db)
		refreshTokenRepository repository.RefreshTokenRepository = repository.NewRefreshTokenRepository(db)

		//=========== (SERVICE) ===========//
		authService service.AuthService = service.NewAuth(userRepository, refreshTokenRepository, mailerService, db)
		// userService                   service.UserService                   = service.NewUser(userRepository, userDisciplineRepository, disciplineGroupConsolidatorRepository, disciplineListDocumentConsolidatorRepository, packageRepository, db)

		//=========== (CONTROLLER) ===========//
		authController controller.AuthController = controller.NewAuth(authService)
		// userController                   controller.UserController                   = controller.NewUser(userService)
	)

	// Register all routes
	routes.Auth(server, authController, middleware)

	return RestConfig{
		server: server,
	}
}

func (ap *RestConfig) Start() {
	port := os.Getenv("APP_PORT")
	host := os.Getenv("APP_HOST")
	if port == "" {
		port = "8080"
	}

	serve := fmt.Sprintf("%s:%s", host, port)
	if err := ap.server.Run(serve); err != nil {
		log.Panicf("failed to start server: %s", err)
	}
	log.Println("server start on port ", serve)
}
