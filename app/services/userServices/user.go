package userServices

import (
	"crypto/sha256"
	"encoding/hex"
	"xinmiao/app/models"
	"xinmiao/config/database"
)

func GetUserByStudentIDAndPassword(userID, password string) (*models.User, error) {
	// 通过学号和密码获取学生用户信息
	h := sha256.New()
	h.Write([]byte(password))
	user := models.User{}
	pass := hex.EncodeToString(h.Sum(nil))
	result := database.DB.Where(
		&models.User{
			UserID:   userID,
			Password: pass,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	DecryptUserKeyInfo(&user)
	return &user, nil
}

func GetUserByID(ID string) (*models.User, error) {
	// 通过学号获取学生用户信息
	user := models.User{}
	result := database.DB.Where(
		&models.User{
			ID: ID,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	DecryptUserKeyInfo(&user)
	return &user, nil
}

func GetUserByUserID(userID string) (*models.User, error) {
	// 通过学号获取学生用户信息
	user := models.User{}
	result := database.DB.Where(
		&models.User{
			UserID: userID,
		},
	).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	DecryptUserKeyInfo(&user)
	return &user, nil
}
