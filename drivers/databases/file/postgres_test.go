package file_test

import (
	"context"
	"database/sql"
	"errors"
	_fileRepo "hungry-baby/drivers/databases/file"
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
	Repository *_fileRepo.PostgresRepository

	DBMock         *gorm.DB
	Mock           sqlmock.Sqlmock
	RepositoryMock *_fileRepo.PostgresRepository
}

var s SQLTest

func SetupSuite(t *testing.T) *sql.DB {
	//SETUP with actual DB
	configApp := _config.GetConfig()
	configDB := _dbDriver.ConfigDB{
		DB_Filename: configApp.Database.File,
		DB_Password: configApp.Database.Pass,
		DB_Host:     configApp.Database.Host,
		DB_Port:     configApp.Database.Port,
		DB_Database: configApp.Database.Name + "_test",
	}

	s.DBConn = configDB.InitialDB()
	s.Repository = _fileRepo.NewPostgresRepository(s.DBConn)

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

	s.RepositoryMock = _fileRepo.NewPostgresRepository(s.DBMock)

	//RETURN dbconnection to close after test
	return db
}

func tearUp(t *testing.T) (func(t *testing.T, db *sql.DB), *sql.DB) {
	//SETUP
	db := SetupSuite(t)
	//MIGRATE
	s.DBConn.AutoMigrate(
		&_fileRepo.Category{},
	)
	//SEED Database
	seeder(s.DBConn)

	return func(t *testing.T, db *sql.DB) {
		//DROP table after test
		s.DBConn.Migrator().DropTable(&_fileRepo.Category{})
		// CLOSE the mock db connection
		db.Close()
	}, db
}

func seeder(db *gorm.DB) {
	var categories = []_fileRepo.Category{
		{
			Title:       "Sport",
			Description: "a pack of sport news",
			Active:      true,
			Archive:     false,
		},
		{
			Title:       "Woman",
			Description: "when woman need a exclusive needs",
			Active:      true,
			Archive:     false,
		},
		{
			Title:       "2019 Indonesian Election",
			Description: "hot list about Election on Indonesia",
			Active:      false,
			Archive:     false,
		},
		{
			Title:       "Terorism",
			Description: "all about terorism in worldwide",
			Active:      true,
			Archive:     true,
		},
	}

	db.Create(&categories)
}

func TestFindByID(t *testing.T) {
	tearDown, db := tearUp(t)
	defer tearDown(t, db)

	t.Run("test case 1 : valid case", func(t *testing.T) {
		id := 1
		result, err := s.Repository.FindByID(id)

		assert.Nil(t, err)
		assert.Equal(t, id, result.ID)
		assert.Equal(t, result.Title, "Sport")
	})

	t.Run("test case 2 : invalid case", func(t *testing.T) {
		result, err := s.Repository.FindByID(10)

		assert.NotNil(t, err)
		assert.Equal(t, 0, result.ID)
	})
}

func TestFind(t *testing.T) {
	tearDown, db := tearUp(t)
	defer tearDown(t, db)

	t.Run("test case 1 : valid case - all data", func(t *testing.T) {
		result, err := s.Repository.Find(context.Background(), "")

		assert.Nil(t, err)
		assert.Equal(t, 3, len(result))
		for _, val := range result {
			assert.NotEqual(t, "Terorism", val.Title)
		}
	})

	t.Run("test case 2 : valid case - active", func(t *testing.T) {
		result, err := s.Repository.Find(context.Background(), "true")
		assert.Nil(t, err)
		assert.Equal(t, 2, len(result))
		for _, val := range result {
			assert.NotEqual(t, "Terorism", val.Title)
		}
	})

	t.Run("test case 3 : valid case - notActive", func(t *testing.T) {
		result, err := s.Repository.Find(context.Background(), "false")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(result))
		assert.Equal(t, "2019 Indonesian Election", result[0].Title)
	})
}

func TestFindWithMock(t *testing.T) {
	tearDown, db := tearUp(t)
	defer tearDown(t, db)

	t.Run("test mock case 1 : invalid case", func(t *testing.T) {
		errorQuery := "mock db error"
		s.Mock.ExpectQuery("SELECT").WithArgs(false, false).WillReturnError(errors.New(errorQuery))

		_, err := s.RepositoryMock.Find(context.Background(), "false")
		assert.NotNil(t, err)
		assert.EqualError(t, err, errorQuery)

		if err := s.Mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expections: %s", err)
		}
	})
}
