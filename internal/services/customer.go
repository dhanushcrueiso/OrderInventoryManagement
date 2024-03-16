package services

import (
	"OrderInventoryManagement/config"
	globals "OrderInventoryManagement/constants"
	"OrderInventoryManagement/internal/database/daos"
	"OrderInventoryManagement/internal/database/models"
	"OrderInventoryManagement/internal/dtos"
	"time"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SaveCustomer(c *fiber.Ctx, req dtos.Customer) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
	req.Password = string(hashedPassword)
	daoReq, err := CustomerDtosToDao(req)
	if err != nil {
		return err
	}

	if err := daos.SaveCustomer(c, daoReq); err != nil {
		return err
	}
	return nil
}

func CustomerLogin(c *fiber.Ctx, req dtos.Customer) (*dtos.LoginRes, error) {
	account, err := daos.GetAccountCustomer(c, req)
	if err != nil {
		return nil, err
	}
	aid := account.ID
	genToken, _ := GetAccessAndRefreshToken(globals.TokenLen)
	token := models.Token{
		Id:         uuid.New(),
		Token:      genToken,
		AccountId:  aid,
		Expires_At: time.Now().Add(time.Duration(config.Get().TokenExpiry) * time.Second),
	}
	err = daos.UpsertToken(token)
	if err != nil {
		return nil, err
	}
	res := dtos.LoginRes{
		ID:       aid,
		Username: account.Username,
		Name:     account.Name,
		Password: "",
		Email:    account.Email,
		Mobile:   account.Mobile,
		Role:     "",
		Token:    genToken,
	}

	return &res, nil
}
func CustomerDtosToDao(req dtos.Customer) (models.Customer, error) {

	return models.Customer{
		ID:       uuid.New(),
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Name:     req.Name,
	}, nil
}
