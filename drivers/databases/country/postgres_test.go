package country_test

import (
	"context"
	"database/sql"
	"errors"
	_countryRepo "hungry-baby/drivers/databases/country"
	"testing"

	_config "hungry-baby/app/config"
	_dbDriver "hungry-baby/drivers/postgres"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SQLTest struct {
	DBConn     *gorm.DB
	Repository *_countryRepo.PostgresRepository

	DBMock         *gorm.DB
	Mock           sqlmock.Sqlmock
	RepositoryMock *_countryRepo.PostgresRepository
}

var s SQLTest

func SetupSuite(t *testing.T) *sql.DB {
	//SETUP with actual DB
	configApp := _config.GetConfig()
	configDB := _dbDriver.ConfigDB{
		DB_Username: configApp.Database.User,
		DB_Password: configApp.Database.Pass,
		DB_Host:     configApp.Database.Host,
		DB_Port:     configApp.Database.Port,
		DB_Database: configApp.Database.Name + "_test",
	}

	s.DBConn = configDB.InitialDB()
	s.Repository = _countryRepo.NewPostgresRepository(s.DBConn)

	//SETUP with mock DB for check the error
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)

	s.Mock = mock
	s.DBMock, err = gorm.Open(
		postgres.New(
			postgres.Config{
				Conn: db,
			},
		),
		&gorm.Config{},
	)
	assert.Nil(t, err)

	s.RepositoryMock = _countryRepo.NewPostgresRepository(s.DBMock)

	//RETURN dbconnection to close after test
	return db
}

func tearUp(t *testing.T) (func(t *testing.T, db *sql.DB), *sql.DB) {
	//SETUP
	db := SetupSuite(t)
	//MIGRATE
	s.DBConn.AutoMigrate(
		&_countryRepo.Country{},
	)
	//SEED Database
	seeder(s.DBConn)

	return func(t *testing.T, db *sql.DB) {
		//DROP table after test
		s.DBConn.Migrator().DropTable(&_countryRepo.Country{})
		// CLOSE the mock db connection
		db.Close()
	}, db
}

func seeder(db *gorm.DB) {
	var categories = []_countryRepo.Country{
		{
			CountryCode: "user_procountry",
			Name:        "user_procountry/1.jpg",
			Status:      true,
		},
		{
			CountryCode: "user_procountry",
			Name:        "user_procountry/2.jpg",
			Status:      true,
		},
		{
			CountryCode: "user_procountry",
			Name:        "user_procountry/3.jpg",
			Status:      true,
		},
		{
			CountryCode: "user_procountry",
			Name:        "user_procountry/4.jpg",
			Status:      true,
		},
	}

	db.Create(&categories)
}

func TestFindByID(t *testing.T) {
	tearDown, db := tearUp(t)
	defer tearDown(t, db)

	t.Run("test case 1 : valid case", func(t *testing.T) {
		id := 1
		result, err := s.Repository.FindByID(context.Background(), id, "true")

		assert.Nil(t, err)
		assert.Equal(t, id, result.ID)
		assert.Equal(t, result.CountryCode, "Sport")
	})

	t.Run("test case 2 : invalid case", func(t *testing.T) {
		result, err := s.Repository.FindByID(context.Background(), 10, "true")

		assert.NotNil(t, err)
		assert.Equal(t, 0, result.ID)
	})
}

func TestFind(t *testing.T) {
	tearDown, db := tearUp(t)
	defer tearDown(t, db)

	t.Run("test case 1 : valid case - all data", func(t *testing.T) {
		result, err := s.Repository.FindAll(context.Background(), "", "true")

		assert.Nil(t, err)
		assert.Equal(t, 3, len(result))
		for _, val := range result {
			assert.NotEqual(t, "Terorism", val.CountryCode)
		}
	})
}

func TestFindWithMock(t *testing.T) {
	tearDown, db := tearUp(t)
	defer tearDown(t, db)

	t.Run("test mock case 1 : invalid case", func(t *testing.T) {
		errorQuery := "mock db error"
		s.Mock.ExpectQuery("SELECT").WithArgs(false, false).WillReturnError(errors.New(errorQuery))

		_, err := s.RepositoryMock.FindAll(context.Background(), "", "true")
		assert.NotNil(t, err)
		assert.EqualError(t, err, errorQuery)

		if err := s.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expections: %s", err)
		}
	})
}
