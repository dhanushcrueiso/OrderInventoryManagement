package daos

import (
	"OrderInventoryManagement/internal/database/db"
	"OrderInventoryManagement/internal/database/models"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UpsertToken(req models.Token) error {
	fmt.Println("check res ", req)
	err := db.DB.Unscoped().Table("tokens").Save(req).Error
	if err != nil {
		return errors.New("failed while upserting token")
	}
	return nil
}

func GetAccountByToken(token string) (*models.Account, error) {
	fmt.Println("check daos")
	res := models.Account{}
	tokenres := models.Token{}
	err := db.DB.Unscoped().Table("tokens").Take(&tokenres).Where("token =?", token).Order("expired_at DESC").Error
	if err != nil {
		return nil, errors.New("failed to get account using tokens")
	}
	fmt.Println("check 1")
	cus := models.Customer{}
	err = db.DB.Unscoped().Table("customers").Where("id =?", tokenres.AccountId).Take(&cus).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.New("failed to get account using tokens")
	}
	fmt.Println("check 2")
	user := models.User{}
	err = db.DB.Unscoped().Table("users").Where("id =?", tokenres.AccountId).Take(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.New("failed to get account using tokens")
	}
	fmt.Println("check 3")
	if cus.ID != uuid.Nil && cus.Username != "" {
		res.Email = cus.Email
		res.Username = cus.Username
		res.ID = cus.ID
	}
	if user.ID != uuid.Nil && user.Username != "" {
		res.Email = user.Email
		res.Username = user.Username
		res.ID = user.ID
	}

	return &res, nil

}
