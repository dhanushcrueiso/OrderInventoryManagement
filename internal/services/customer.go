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

func PlaceOrder(c *fiber.Ctx, req []dtos.Product, customerId uuid.UUID) (string, error) {
	orderReq := []models.Order{}
	inventoryUpdate := []models.Inventory{}
	Id := uuid.New()
	orderId, _ := GetAccessAndRefreshToken(globals.OrderIdLen)
	for _, product := range req {
		order := models.Order{
			CustomerID:      customerId,
			ID:              Id,
			ProductID:       product.ID,
			QuantityOrdered: product.Quantity,
			TotalPrice:      product.Price * float64(product.Quantity),
			OrderId:         orderId,
			OrderDate:       time.Now(),
		}
		orderReq = append(orderReq, order)

		invent := models.Inventory{
			ProductID:   product.ID,
			Quantity:    product.Quantity,
			LastUpdated: time.Now(),
		}
		inventoryUpdate = append(inventoryUpdate, invent)
	}

	err := daos.UpsertOrders(orderReq)
	if err != nil {
		return "", err
	}
	for _, inv := range inventoryUpdate {
		err := daos.UpdateInventory(c, inv, "subtract")
		if err != nil {
			return "", err
		}
	}
	return orderId, nil
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
