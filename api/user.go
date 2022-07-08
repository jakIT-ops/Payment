package api

import (
	"log"
	"payment_full/models"
	"payment_full/services"
	"payment_full/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

var serviceUser services.User

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user models.User) userResponse {
	return userResponse{
		Username:          user.UserName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *fiber.Ctx) error {
	var req createUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	log.Println("before hash pass:", req.Password)
	hashedPassword, err := utils.HashPassword(req.Password)
	log.Println("after hash pass:", hashedPassword)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	arg := services.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
		Email:    req.Email,
	}

	user, err := serviceUser.CreateUser(arg)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}

	rsp := newUserResponse(user)
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"status":  "Амжилттай хадаглалаа",
		"account": rsp,
	})
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

// type loginUserResponse struct {
// 	SessionID             uuid.UUID    `json:"session_id"`
// 	AccessToken           string       `json:"access_token"`
// 	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
// 	RefreshToken          string       `json:"refresh_token"`
// 	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
// 	User                  userResponse `json:"user"`
// }
type loginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *fiber.Ctx) error {
	var req loginUserRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	log.Println("test1")

	user, err := serviceUser.GetUser(req.Username)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
	}
	log.Println("test2", user)

	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	log.Println("test3")

	accessToken, err := server.tokenMaker.CreateToken(
		user.UserName,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"error":   err.Error(),
		})
	}
	log.Println("test4")

	rsp := &loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": "Амжилттай нэвтэрлэлээ",
		"data":    rsp,
	})
}

// session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
// 	ID:           refreshPayload.ID,
// 	Username:     user.Username,
// 	RefreshToken: refreshToken,
// 	UserAgent:    ctx.Request.UserAgent(),
// 	ClientIp:     ctx.ClientIP(),
// 	IsBlocked:    false,
// 	ExpiresAt:    refreshPayload.ExpiredAt,
// })
// if err != nil {
// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 	return
// }

// rsp := loginUserResponse{
// 	SessionID:             session.ID,
// 	AccessToken:           accessToken,
// 	AccessTokenExpiresAt:  accessPayload.ExpiredAt,
// 	RefreshToken:          refreshToken,
// 	RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
// 	User:                  newUserResponse(user),
// }
// ctx.JSON(http.StatusOK, rsp)
