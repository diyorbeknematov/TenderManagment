package postgres_test

// import (
// 	"tender/config"
// 	"tender/logs"
// 	"tender/model"
// 	"tender/storage"
// 	"tender/storage/postgres"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateUser(t *testing.T) {
// 	db := postgres.ConnectDB(postgres.Logger)

// 	userRepo := storage.NewStorage(db, logs.InitLogger())

// 	resp, err := userRepo.RegistrationRepository().CreateUser(model.UserRegisterReq{
// 		Username: "diyorbeknematov1",
// 		Email:    "diyorbeknematov1@gmail.com",
// 		Role:     "client",
// 		Password: "123456789",
// 	})

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.NotNil(t, resp)
// 	assert.Equal(t, "User registered successfully", resp.Message)
// }

// func TestGetUserByEmail(t *testing.T) {
// 	db, err := postgres.Connect(config.LoadConfig())
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	UserRepo := storage.NewStorage(db, logs.InitLogger())

// 	resp, err := UserRepo.RegistrationRepository().GetUserByEmail("diyorbeknematov@gmail.com")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.NotNil(t, resp)
// }

// func TestIsUserExists(t *testing.T) {
// 	db, err := postgres.Connect(config.LoadConfig())
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	UserRepo := storage.NewStorage(db, logs.InitLogger())

// 	exists, err := UserRepo.RegistrationRepository().IsUserExists("diyorbeknematov@gmail.com", "diyorbeknematov")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	assert.Equal(t, true, exists)
// }
