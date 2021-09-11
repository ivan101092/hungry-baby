package main

import (
	"fmt"
	_dbFactory "hungry-baby/drivers/databases"

	_fileUsecase "hungry-baby/businesses/file"
	_fileController "hungry-baby/controllers/file"

	_countryUsecase "hungry-baby/businesses/country"
	_countryController "hungry-baby/controllers/country"

	_provinceUsecase "hungry-baby/businesses/province"
	_provinceController "hungry-baby/controllers/province"

	_cityUsecase "hungry-baby/businesses/city"
	_cityController "hungry-baby/controllers/city"

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
	db.AutoMigrate()
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

	fileRepo := _dbFactory.NewFileRepository(db)
	fileUsecase := _fileUsecase.NewFileUsecase(timeoutContext, fileRepo, connMinio)
	fileCtrl := _fileController.NewFileController(fileUsecase)

	countryRepo := _dbFactory.NewCountryRepository(db)
	countryUsecase := _countryUsecase.NewCountryUsecase(timeoutContext, countryRepo)
	countryCtrl := _countryController.NewCountryController(countryUsecase)

	provinceRepo := _dbFactory.NewProvinceRepository(db)
	provinceUsecase := _provinceUsecase.NewProvinceUsecase(timeoutContext, provinceRepo)
	provinceCtrl := _provinceController.NewProvinceController(provinceUsecase)

	cityRepo := _dbFactory.NewCityRepository(db)
	cityUsecase := _cityUsecase.NewCityUsecase(timeoutContext, cityRepo)
	cityCtrl := _cityController.NewCityController(cityUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:      configJWT.Init(),
		FileController:     *fileCtrl,
		CountryController:  *countryCtrl,
		ProvinceController: *provinceCtrl,
		CityController:     *cityCtrl,
	}
	routesInit.RouteRegister(e)

	fmt.Println(configJWT.GenerateToken(1))

	log.Fatal(e.Start(viper.GetString("server.address")))
}
