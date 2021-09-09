package file_test

import (
	"context"
	"database/sql"
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
		DB_Username: configApp.Database.User,
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
		&_fileRepo.File{},
	)
	//SEED Database
	seeder(s.DBConn)

	return func(t *testing.T, db *sql.DB) {
		//DROP table after test
		s.DBConn.Migrator().DropTable(&_fileRepo.File{})
		// CLOSE the mock db connection
		db.Close()
	}, db
}

func seeder(db *gorm.DB) {
	var categories = []_fileRepo.File{
		{
			Type:       "user_profile",
			URL:        "user_profile/1.jpg",
			UserUpload: "1",
		},
		{
			Type:       "user_profile",
			URL:        "user_profile/2.jpg",
			UserUpload: "1",
		},
		{
			Type:       "user_profile",
			URL:        "user_profile/3.jpg",
			UserUpload: "1",
		},
		{
			Type:       "user_profile",
			URL:        "user_profile/4.jpg",
			UserUpload: "",
		},
	}

	db.Create(&categories)
}

func TestFindByID(t *testing.T) {
	tearDown, db := tearUp(t)
	defer tearDown(t, db)

	t.Run("test case 1 : valid case", func(t *testing.T) {
		id := 1
		result, err := s.Repository.FindByID(context.Background(), id)

		assert.Nil(t, err)
		assert.Equal(t, id, result.ID)
		assert.Equal(t, result.Type, "Sport")
	})

	t.Run("test case 2 : invalid case", func(t *testing.T) {
		result, err := s.Repository.FindByID(context.Background(), 10)

		assert.NotNil(t, err)
		assert.Equal(t, 0, result.ID)
	})
}
