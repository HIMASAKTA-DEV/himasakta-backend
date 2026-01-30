package config

import (
	"fmt"

	"log"
	"os"

	"github.com/azkaazkun/be-samarta/db"
	"github.com/azkaazkun/be-samarta/internal/api/controller"
	"github.com/azkaazkun/be-samarta/internal/api/repository"
	"github.com/azkaazkun/be-samarta/internal/api/routes"
	"github.com/azkaazkun/be-samarta/internal/api/service"
	"github.com/azkaazkun/be-samarta/internal/middleware"
	mailer "github.com/azkaazkun/be-samarta/internal/pkg/email"
	"github.com/azkaazkun/be-samarta/internal/pkg/google/oauth"
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
		oauthService  oauth.Oauth   = oauth.New()
		// awsS3Service  storage.AwsS3 = storage.NewAwsS3()

		//=========== (REPOSITORY) ===========//
		userRepository         repository.UserRepository         = repository.NewUser(db)
		refreshTokenRepository repository.RefreshTokenRepository = repository.NewRefreshTokenRepository(db)
		sptjmRepository        repository.SPTJMRepository        = repository.NewSPTJM(db)
		sshRepository          repository.SSHRepository          = repository.NewSSH(db)
		itemRepository         repository.ItemRepository         = repository.NewItem(db)
		proposalRepository     repository.ProposalRepository     = repository.NewProposal(db)

		//=========== (SERVICE) ===========//
		authService     service.AuthService     = service.NewAuth(userRepository, refreshTokenRepository, mailerService, oauthService, db)
		sptjmService    service.SPTJMService    = service.NewSPTJM(sptjmRepository, db)
		sshService      service.SSHService      = service.NewSSH(sshRepository, itemRepository, proposalRepository, db)
		itemService     service.ItemService     = service.NewItem(itemRepository, db)
		proposalService service.ProposalService = service.NewProposal(proposalRepository, db)
		// userService                   service.UserService                   = service.NewUser(userRepository, userDisciplineRepository, disciplineGroupConsolidatorRepository, disciplineListDocumentConsolidatorRepository, packageRepository, db)

		//=========== (CONTROLLER) ===========//
		authController     controller.AuthController     = controller.NewAuth(authService)
		sptjmController    controller.SPTJMController    = controller.NewSPTJM(sptjmService)
		sshController      controller.SSHController      = controller.NewSSH(sshService)
		itemController     controller.ItemController     = controller.NewItem(itemService)
		proposalController controller.ProposalController = controller.NewProposal(proposalService)
		// userController                   controller.UserController                   = controller.NewUser(userService)
	)

	// Register all routes
	routes.Auth(server, authController, middleware)
	routes.SPTJM(server, sptjmController, middleware)
	routes.SSH(server, sshController, middleware)
	routes.Item(server, itemController, middleware)
	routes.Proposal(server, proposalController, middleware)

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
