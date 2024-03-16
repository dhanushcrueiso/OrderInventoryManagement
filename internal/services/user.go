package services

import (
	"OrderInventoryManagement/config"
	globals "OrderInventoryManagement/constants"
	"OrderInventoryManagement/internal/database/daos"
	"OrderInventoryManagement/internal/dtos"
	"fmt"
	"time"

	"OrderInventoryManagement/internal/database/models"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CheckName(c *fiber.Ctx, req string) error {

	res, err := daos.GetName(c, req)
	if err != nil {
		return err
	}
	fmt.Println("res", res)
	return nil
}

func SaveUser(c *fiber.Ctx, req dtos.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil
	}
	req.Password = string(hashedPassword)
	daoReq, err := DtosToDao(req)
	if err != nil {
		return err
	}

	if err := daos.SaveUser(c, daoReq); err != nil {
		return err
	}
	return nil
}

func Login(c *fiber.Ctx, req dtos.User) (*dtos.LoginRes, error) {

	account, err := daos.GetAccount(c, req)
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
		Mobile:   account.Email,
		Role:     account.Role,
		Token:    genToken,
	}
	return &res, nil

}

func DtosToDao(req dtos.User) (models.User, error) {

	return models.User{
		ID:       uuid.New(),
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
		Role:     globals.RoleExecutive,
		Mobile:   req.Mobile,
		Name:     req.Name,
	}, nil
}
