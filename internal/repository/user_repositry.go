// user_repository.go
package repository

import (
	"errors"
	"fmt"
	"go-project-api/internal/model"
	"go-project-api/internal/security"
	"go-project-api/internal/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

// this is global access to the functions in this page... I know this is not the correct way to do this
// But I am struggleing to figure this out for now... I need access to a DB instane inorder to use any of these functions
// and this is proving rather difficult from distant packages such as middleware.
var UserRepoInstance *UserRepository

func GetUserByID(r *UserRepository, id int) (*model.User, error) {
	var user model.User
	result := r.DB.First(&user, id)
	if result.Error != nil {
		util.WriteLog("Attempting to get user with ID:", id)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID %d not found", id)
		} else if result.Error != nil {
			return nil, result.Error
		}
		return nil, result.Error

	}

	return &user, nil
}

func UpdateUser(r *UserRepository, user *model.User) error {
	result := r.DB.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteUser(r *UserRepository, id uint) error {
	result := r.DB.Delete(&model.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func CreateUser(r *UserRepository, user *model.User) error {

	// Hash the password before saving the user
	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	result := r.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindByUsername(r *UserRepository, username string) (*model.User, error) {
	// Implement the logic to find a user by username in the database
	// Example using a hypothetical ORM:
	var user model.User
	result := r.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func UsernameExists(r *UserRepository, username string) (bool, error) {
	// Implement the logic to check if a username already exists in the database

	var count int64
	result := r.DB.Model(&model.User{}).Where("username = ?", username).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func ConnectDatabase() (*UserRepository, error) {

	//if UserRepoInstance != nil {
	//	return UserRepoInstance, nil
	//}

	dsn := "rooty_tooty:Blast123@tcp(localhost:3306)/users_database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to the database.")
		return nil, err
	}

	// Auto-migrate the User schema
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Println("Error auto-migrating the User schema.")
		return nil, err
	}

	log.Printf("Successfully connected to database. ")

	userRepo := NewUserRepository(db)

	//UserRepoInstance = &UserRepository{DB: db}
	return userRepo, nil
}

func (r *UserRepository) Close() error {
	sqlDB, err := r.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
