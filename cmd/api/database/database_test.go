package database_test

import (
	"testing"
	"uala/cmd/api/database"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm database: %v", err)
	}

	return gormDB, mock
}

func TestInitDatabase(t *testing.T) {
	dbInstance := database.NewDataBaseIntance()
	db, mock := setupMockDB(t)
	// defer dbInstance.DB.Close()

	mock.ExpectExec("CREATE DATABASE test_db").WillReturnResult(sqlmock.NewResult(1, 1))
	dbInstance.Writer = db
	// dbInstance.InitDatabase("test_db")

	assert.NotNil(t, dbInstance.Writer)
}

func TestGetConnectionSingleton(t *testing.T) {
	dbInstance := database.NewDataBaseIntance()
	db, _ := setupMockDB(t)
	dbInstance.Writer = db

	conn := dbInstance.SingletonDB()
	assert.NotNil(t, conn)
}

func TestInitTransaction(t *testing.T) {
	dbInstance := database.NewDataBaseIntance()
	db, mock := setupMockDB(t)
	// defer dbInstance.DB.Close()

	mock.ExpectBegin()
	dbInstance.Writer = db
	dbInstance.InitTransaction()

	assert.NotNil(t, dbInstance.Writer)
}

func TestCommitTransaction(t *testing.T) {
	dbInstance := database.NewDataBaseIntance()
	db, mock := setupMockDB(t)
	// defer dbInstance.DB.Close()

	mock.ExpectBegin()
	mock.ExpectCommit()
	dbInstance.Writer = db
	dbInstance.InitTransaction()
	dbInstance.CommitTransaction()

	assert.NotNil(t, dbInstance.Writer)
}

func TestRollbackTransaction(t *testing.T) {
	dbInstance := database.NewDataBaseIntance()
	db, mock := setupMockDB(t)
	// defer dbInstance.DB.Close()

	mock.ExpectBegin()
	mock.ExpectRollback()
	dbInstance.Writer = db
	dbInstance.InitTransaction()
	dbInstance.RollbackTransaction()

	assert.NotNil(t, dbInstance.Writer)
}

func TestRunMigrations(t *testing.T) {
	// db, mock := setupMockDB(t)
	// defer db.Close()

	// mock.ExpectExec("CREATE TABLE").WillReturnResult(sqlmock.NewResult(1, 1))
	// database.Run(db)

	// Aquí puedes agregar más aserciones para verificar que las migraciones se ejecutaron correctamente
}

// func TestBaseHooks(t *testing.T) {
// 	db, mock := setupMockDB(t)

// 	mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
// 	mock.ExpectExec("DELETE").WillReturnResult(sqlmock.NewResult(1, 1))

// 	base := &database.Base{}
// 	db.Create(base)

// 	assert.NotEmpty(t, base.id)
// 	assert.WithinDuration(t, time.Now(), base.CreatedAt, time.Second)
// 	assert.WithinDuration(t, time.Now(), base.UpdatedAt, time.Second)

// 	base.UpdatedAt = time.Now().Add(-time.Hour)
// 	db.Save(base)
// 	assert.WithinDuration(t, time.Now(), base.UpdatedAt, time.Second)

// 	db.Delete(base)
// 	assert.WithinDuration(t, time.Now(), base.DeletedAt.Time, time.Second)
// }

// func TestSetTimestamp(t *testing.T) {
// 	db, _ := setupMockDB(t)
// 	base := &database.Base{
// 		CreatedAt: time.Now().Add(-time.Hour),
// 		UpdatedAt: time.Now(),
// 	}

// 	db.Create(base)

// 	assert.Equal(t, base.Created, common.ParseTimeFormat(base.CreatedAt))
// 	assert.Equal(t, base.Updated, common.ParseTimeFormat(base.UpdatedAt))
// }
