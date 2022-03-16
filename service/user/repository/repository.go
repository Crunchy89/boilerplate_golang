package repository

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/Crunchy89/boilerplate_golang/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (user *entities.User) BeforeSave() error {
	hashedPassword, err := Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *entities.User) Prepare() {
	user.ID = 0
	user.Nickname = html.EscapeString(strings.TrimSpace(user.Nickname))
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}
func (user *entities.User) SaveUser(db *gorm.DB) (*entities.User, error) {

	var err error
	err = db.Debug().Create(&user).Error
	if err != nil {
		return &entities.User{}, err
	}
	return user, nil
}

func (user *entities.User) FindAllUsers(db *gorm.DB) (*[]entities.User, error) {
	var err error
	users := []entities.User{}
	err = db.Debug().Model(&entities.User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]entities.User{}, err
	}
	return &users, err
}

func (user *entities.User) FindUserByID(db *gorm.DB, uid uint32) (*entities.User, error) {
	var err error
	err = db.Debug().Model(entities.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &entities.User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &entities.User{}, errors.New("User Not Found")
	}
	return user, err
}

func (user *entities.User) UpdateAUser(db *gorm.DB, uid uint32) (*entities.User, error) {

	// To hash the password
	err := user.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&entities.User{}).Where("id = ?", uid).Take(&entities.User{}).UpdateColumns(
		map[string]interface{}{
			"password":  user.Password,
			"nickname":  user.Nickname,
			"email":     user.Email,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &entities.User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&entities.User{}).Where("id = ?", uid).Take(&user).Error
	if err != nil {
		return &entities.User{}, err
	}
	return user, nil
}

func (user *entities.User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&entities.User{}).Where("id = ?", uid).Take(&entities.User{}).Delete(&entities.User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
