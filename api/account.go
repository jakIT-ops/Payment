package api

import (
	"database/sql"
	"errors"
	"net/http"
	db "payment_full/db/sqlc"
	"payment_full/token"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createAccountRequest struct {
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateAccountParams{
		Owner:    authPayload.Username,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if account.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}

//var service services.Account
//
//type AccountServer struct{}

//
//type createAccountRequest struct {
//	Owner    string  `json:"owner"`
//	Currency string  `json:"currency"`
//	Balance  float64 `json:"amount"`
//}
//
//func (a *AccountServer) createAccount(ctx *fiber.Ctx) error {
//	var req createAccountRequest
//	if err := ctx.BodyParser(&req); err != nil {
//		return ctx.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	}
//
//	arg := services.CreateAccountParams{
//		Owner:    req.Owner,
//		Currency: req.Currency,
//		Balance:  0,
//	}
//
//	account, err := service.CreateAccount(arg)
//	if err != nil {
//		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	}
//	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
//		"account": account,
//		"status":  "aмжилттай хадаглалаа",
//	})
//}
//
//type getAccountRequest struct {
//	ID string `json:"id"`
//}
//
//func (a *AccountServer) getAccount(ctx *fiber.Ctx) error {
//	var req getAccountRequest
//	req.ID = ctx.Params("id")
//
//	account, err := service.GetAccount(req.ID)
//	if err != nil {
//		return ctx.Status(400).JSON(&fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	}
//
//	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
//		"success": true,
//		"account": account,
//	})
//}
//
//func (a *AccountServer) listAccounts(ctx *fiber.Ctx) error {
//	accounts, err := service.ListAccounts()
//	if err != nil {
//		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
//			"success": false,
//			"error ":  err.Error(),
//		})
//	}
//
//	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
//		"success":  true,
//		"accounts": accounts,
//	})
//}
//
//type updateAccountReq struct {
//	Id      string `json:"id"`
//	Balance int64  `json:"amount"`
//}
//
//func (a *AccountServer) updateAccount(ctx *fiber.Ctx) error {
//	var req updateAccountReq
//
//	if err := ctx.BodyParser(&req); err != nil {
//		return ctx.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	}
//
//	arg := services.UpdateAccountParams{
//		Id:      req.Id,
//		Balance: req.Balance,
//	}
//
//	account, err := service.UpdateAccount(arg)
//	if err != nil {
//		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	}
//	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
//		"account": account,
//		"status":  "амжилттай шинэчиллээ",
//	})
//}
//
//func (a *AccountServer) deleteAccount(ctx *fiber.Ctx) error {
//	id := ctx.Params("id")
//
//	if err := service.DeleteAccount(id); err != nil {
//		return ctx.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	}
//	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
//		"success": true,
//		"status":  "Амжилттай устгаглаа",
//	})
//}
