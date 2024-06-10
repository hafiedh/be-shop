package main

import (
	"be-shop/internal/app"
	"be-shop/internal/app/controller"
	"be-shop/internal/app/infra"
	"be-shop/internal/app/repo/postgres"
	"be-shop/internal/app/service"
	"be-shop/internal/app/service/utils"
	"be-shop/pkg/di"
	"be-shop/pkg/middleware"
	"fmt"
	"log/slog"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file")
	}

	utils.InitValidator()

	err = infra.InitTimezone()
	if err != nil {
		slog.Error("Error loading timezone")
	}

	err = LoadApplicationConfig()
	if err != nil {
		fmt.Println("LoadApplicationConfig: ", err.Error())
		slog.Error(err.Error())
	}

	err = LoadApplicationPackage()
	if err != nil {
		fmt.Println("LoadApplicationPackage: ", err.Error())
		slog.Error(err.Error())
	}

	err = LoadApplicationRepository()
	if err != nil {
		fmt.Println("LoadApplicationRepository: ", err.Error())
		slog.Error(err.Error())
	}

	err = LoadApplicationService()
	if err != nil {
		fmt.Println("LoadApplicationService: ", err.Error())
		slog.Error(err.Error())
	}

	err = LoadApplicationController()
	if err != nil {
		fmt.Println("LoadApplicationController: ", err.Error())
		slog.Error(err.Error())
	}
	app.Start()
}

func LoadApplicationConfig() error {
	err := di.Provide(infra.LoadPgDatabaseCfg)
	if err != nil {
		return fmt.Errorf("LoadPgDatabaseCfg: %s", err.Error())
	}

	err = di.Provide(infra.LoadAppCfg)
	if err != nil {
		return fmt.Errorf("LoadAppCfg: %s", err.Error())
	}

	err = di.Provide(infra.LoadJwtCfg)
	if err != nil {
		return fmt.Errorf("LoadJwtCfg: %s", err.Error())
	}
	return nil
}

func LoadApplicationPackage() error {
	err := di.Provide(infra.NewEcho)
	if err != nil {
		return fmt.Errorf("NewEcho: %s", err.Error())
	}

	err = di.Provide(infra.NewDatabases)
	if err != nil {
		fmt.Println("NewDatabases: ", err.Error())
		return fmt.Errorf("NewDatabases: %s", err.Error())
	}
	return nil
}

func LoadApplicationRepository() error {
	err := di.Provide(postgres.NewUserRepo)
	if err != nil {
		return fmt.Errorf("NewMoneyTransferRepo: %s", err.Error())
	}
	err = di.Provide(postgres.NewProductRepo)
	if err != nil {
		return fmt.Errorf("NewProductRepo: %s", err.Error())
	}
	return nil
}

func LoadApplicationService() error {
	err := di.Provide(service.NewUserSvc)
	if err != nil {
		return fmt.Errorf("NewUserSvc: %s", err.Error())
	}

	err = di.Provide(service.NewProductSvc)
	if err != nil {
		return fmt.Errorf("NewProductSvc: %s", err.Error())
	}

	return nil
}

func LoadApplicationController() error {
	err := di.Provide(controller.NewAuthCtrl)
	if err != nil {
		return fmt.Errorf("NewAuthCtrl: %s", err.Error())
	}
	err = di.Provide(middleware.NewMiddleWare)
	if err != nil {
		return fmt.Errorf("NewMiddleWare: %s", err.Error())
	}

	err = di.Provide(controller.NewProductCtrl)
	if err != nil {
		return fmt.Errorf("NewProductCtrl: %s", err.Error())
	}

	return nil
}
