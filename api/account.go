package api

import (
	"net/http"
	"payment_full/services"

	"github.com/gofiber/fiber/v2"
)

var service services.Account

type AccountServer struct{}

type createAccountRequest struct {
	Owner    string  `json:"owner"`
	Currency string  `json:"currency"`
	Balance  float64 `json:"amount"`
}

func (a *AccountServer) createAccount(ctx *fiber.Ctx) error {
	var req createAccountRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	arg := services.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := service.CreateAccount(arg)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"account": account,
		"status":  "aмжилттай хадаглалаа",
	})
}

type getAccountRequest struct {
	ID string `json:"id"`
}

func (a *AccountServer) getAccount(ctx *fiber.Ctx) error {
	var req getAccountRequest
	req.ID = ctx.Params("id")

	account, err := service.GetAccount(req.ID)
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"account": account,
	})
}

func (a *AccountServer) listAccounts(ctx *fiber.Ctx) error {
	accounts, err := service.ListAccounts()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error ":  err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"success":  true,
		"accounts": accounts,
	})
}

type updateAccountReq struct {
	Id      string `json:"id"`
	Balance int64  `json:"amount"`
}

func (a *AccountServer) updateAccount(ctx *fiber.Ctx) error {
	var req updateAccountReq

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	arg := services.UpdateAccountParams{
		Id:      req.Id,
		Balance: req.Balance,
	}

	account, err := service.UpdateAccount(arg)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"account": account,
		"status":  "амжилттай шинэчиллээ",
	})
}

func (a *AccountServer) deleteAccount(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if err := service.DeleteAccount(id); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"status":  "Амжилттай устгаглаа",
	})
}
