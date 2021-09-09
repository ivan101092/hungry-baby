package main

import (
	_dbFactory "hungry-baby/drivers/databases"

	_countryUsecase "hungry-baby/businesses/country"
	_countryController "hungry-baby/controllers/country"
	_countryRepo "hungry-baby/drivers/databases/country"

	_fileUsecase "hungry-baby/businesses/file"
	_fileController "hungry-baby/controllers/file"

	_minio "hungry-baby/drivers/minio"
	_dbDriver "hungry-baby/drivers/postgres"

	_config "hungry-baby/app/config"
	_middleware "hungry-baby/app/middleware"
	_routes "hungry-baby/app/routes"

	"log"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_countryRepo.Country{},
	)
}

func main() {
	configApp := _config.GetConfig()
	configDB := _dbDriver.ConfigDB{
		DB_Username: configApp.Database.User,
		DB_Password: configApp.Database.Pass,
		DB_Host:     configApp.Database.Host,
		DB_Port:     configApp.Database.Port,
		DB_Database: configApp.Database.Name,
	}
	db := configDB.InitialDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	configMinio := _minio.Connection{
		AccessKey: configApp.Minio.AccessKey,
		SecretKey: configApp.Minio.SecretKey,
		UseSSL:    configApp.Minio.UseSSL,
		BaseURL:   configApp.Minio.Host,
		Duration:  configApp.Minio.Duration,
		Bucket:    configApp.Minio.DefaultBucket,
	}
	minioClient, err := configMinio.InitClient()
	if err != nil {
		panic(err)
	}
	connMinio := _minio.NewMinioModel(minioClient, configMinio.Bucket)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	countryRepo := _dbFactory.NewCountryRepository(db)
	countryUsecase := _countryUsecase.NewCountryUsecase(timeoutContext, countryRepo)
	countryCtrl := _countryController.NewCountryController(countryUsecase)

	fileRepo := _dbFactory.NewFileRepository(db)
	fileUsecase := _fileUsecase.NewFileUsecase(timeoutContext, fileRepo, connMinio)
	fileCtrl := _fileController.NewFileController(fileUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:     configJWT.Init(),
		FileController:    *fileCtrl,
		CountryController: *countryCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
