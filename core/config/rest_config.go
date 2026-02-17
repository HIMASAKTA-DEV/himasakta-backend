package config

import (
	"fmt"

	"log"
	"os"

	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/controller"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/repository"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/routes"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/api/service"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/middleware"
	myjwt "github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/jwt"
	"github.com/HIMASAKTA-DEV/himasakta-backend/core/pkg/storage"
	"github.com/HIMASAKTA-DEV/himasakta-backend/db"
	"github.com/gin-gonic/gin"
)

type RestConfig struct {
	server *gin.Engine
}

func NewRest() (RestConfig, error) {
	db := db.New()
	if db == nil {
		return RestConfig{}, fmt.Errorf("database connection failed")
	}
	app := gin.Default()
	s3 := storage.NewAwsS3()
	server := NewRouter(app, s3)
	_ = middleware.New(db)

	// Injections
	galleryRepo := repository.NewGallery(db)
	galleryService := service.NewGallery(galleryRepo)
	galleryController := controller.NewGallery(galleryService, s3)

	deptRepo := repository.NewDepartment(db)
	deptService := service.NewDepartment(deptRepo)
	deptController := controller.NewDepartment(deptService)

	cabinetRepo := repository.NewCabinetInfo(db)
	cabinetService := service.NewCabinetInfo(cabinetRepo)
	cabinetController := controller.NewCabinetInfo(cabinetService)

	memberRepo := repository.NewMember(db)
	memberService := service.NewMember(memberRepo)
	memberController := controller.NewMember(memberService)

	progendaRepo := repository.NewProgenda(db)
	timelineRepo := repository.NewTimeline(db)
	progendaService := service.NewProgenda(db, progendaRepo, timelineRepo)
	progendaController := controller.NewProgenda(progendaService)

	monthlyEventRepo := repository.NewMonthlyEvent(db)
	monthlyEventService := service.NewMonthlyEvent(monthlyEventRepo)
	monthlyEventController := controller.NewMonthlyEvent(monthlyEventService)

	newsRepo := repository.NewNews(db)
	newsService := service.NewNews(newsRepo)
	newsController := controller.NewNews(newsService)

	nrpWhitelistRepo := repository.NewNrpWhitelist(db)
	nrpWhitelistService := service.NewNrpWhitelist(nrpWhitelistRepo)
	nrpWhitelistController := controller.NewNrpWhitelist(nrpWhitelistService)

	jwtService := myjwt.NewJWT()
	authService := service.NewAuth(jwtService)
	authController := controller.NewAuth(authService)

	indexController := controller.NewIndex()

	roleRepo := repository.NewRole(db)
	roleService := service.NewRole(roleRepo)
	roleController := controller.NewRole(roleService)

	m := middleware.New(db)

	// Register all routes
	server.GET("/", indexController.Index)
	routes.Auth(server, authController)
	routes.Gallery(server, galleryController, m)
	routes.Department(server, deptController, m)
	routes.CabinetInfo(server, cabinetController, m)
	routes.Member(server, memberController, m)
	routes.Progenda(server, progendaController, m)
	routes.MonthlyEvent(server, monthlyEventController, m)
	routes.News(server, newsController, m)
	routes.NrpWhitelist(server, nrpWhitelistController, m)
	routes.Role(server, roleController, m)

	return RestConfig{
		server: server,
	}, nil
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

func (ap *RestConfig) GetServer() *gin.Engine {
	return ap.server
}
