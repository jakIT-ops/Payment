package api

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	"net/http"
	db "payment_full/db/sqlc"
	"payment_full/utils"
	"time"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	SessionID             uuid.UUID    `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  userResponse `json:"user"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = utils.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)
}

//
//var serviceUser services.User
//
//type createUserRequest struct {
//	Username string `json:"username" binding:"required,alphanum"`
//	Password string `json:"password" binding:"required,min=6"`
//	Email    string `json:"email" binding:"required,email"`
//}
//
//type userResponse struct {
//	Username          string    `json:"username"`
//	Email             string    `json:"email"`
//	PasswordChangedAt time.Time `json:"password_changed_at"`
//	CreatedAt         time.Time `json:"created_at"`
//}
//
//func newUserResponse(user models.User) userResponse {
//	return userResponse{
//		Username:          user.UserName,
//		Email:             user.Email,
//		PasswordChangedAt: user.PasswordChangedAt,
//		CreatedAt:         user.CreatedAt,
//	}
//}
//
//func CreateUsers(ctx *fiber.Ctx) error {
//	var req createUserRequest
//	if err := ctx.BodyParser(&req); err != nil {
//		return ctx.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	}
//	log.Println("before hash pass:", req.Password)
//	hashedPassword, err := utils.HashPassword(req.Password)
//	log.Println("after hash pass:", hashedPassword)
//	if err != nil {
//		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	}
//
//	arg := services.CreateUserParams{
//		Username: req.Username,
//		Password: hashedPassword,
//		Email:    req.Email,
//	}
//
//	user, err := serviceUser.CreateUser(arg)
//	if err != nil {
//		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	}
//
//	rsp := newUserResponse(user)
//	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
//		"status":  "Амжилттай хадаглалаа",
//		"account": rsp,
//	})
//}
//
//func (server *Server) createUser(ctx *fiber.Ctx) error {
//	var req createUserRequest
//	if err := ctx.BodyParser(&req); err != nil {
//		return ctx.Status(fiber.ErrBadRequest.Code).JSON(&fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	}
//	log.Println("before hash pass:", req.Password)
//	hashedPassword, err := utils.HashPassword(req.Password)
//	log.Println("after hash pass:", hashedPassword)
//	if err != nil {
//		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	}
//
//	arg := services.CreateUserParams{
//		Username: req.Username,
//		Password: hashedPassword,
//		Email:    req.Email,
//	}
//
//	user, err := serviceUser.CreateUser(arg)
//	if err != nil {
//		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	}
//
//	rsp := newUserResponse(user)
//	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
//		"status":  "Амжилттай хадаглалаа",
//		"account": rsp,
//	})
//}
//
//type loginUserRequest struct {
//	Username string `json:"username" binding:"required,alphanum"`
//	Password string `json:"password" binding:"required,min=6"`
//}
//
////	type loginUserResponse struct {
////		SessionID             uuid.UUID    `json:"session_id"`
////		AccessToken           string       `json:"access_token"`
////		AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
////		RefreshToken          string       `json:"refresh_token"`
////		RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
////		User                  userResponse `json:"user"`
////	}
//type loginUserResponse struct {
//	AccessToken string       `json:"access_token"`
//	User        userResponse `json:"user"`
//}
//
//func (server *Server) loginUser(ctx *fiber.Ctx) error {
//	var req loginUserRequest
//	if err := ctx.BodyParser(&req); err != nil {
//		return ctx.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	}
//	log.Println("test1")
//
//	user, err := serviceUser.GetUser(req.Username)
//	if err != nil {
//		return ctx.Status(fiber.StatusNotFound).JSON(&fiber.Map{
//			"success": false,
//			"error":   err,
//		})
//	}
//	log.Println("test2", user)
//
//	err = utils.CheckPassword(req.Password, user.Password)
//	if err != nil {
//		return ctx.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	}
//	log.Println("test3")
//
//	//accessToken, err := server.tokenMaker.CreateToken(
//	//	user.UserName,
//	//	server.config.AccessTokenDuration,
//	//)
//	if err != nil {
//		return ctx.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
//			"success": false,
//			"error":   err.Error(),
//		})
//	}
//	log.Println("test4")
//
//	rsp := &loginUserResponse{
//		//AccessToken: accessToken,
//		User: newUserResponse(user),
//	}
//	return ctx.Status(fiber.StatusOK).JSON(&fiber.Map{
//		"success": "Амжилттай нэвтэрлэлээ",
//		"data":    rsp,
//	})
//}

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
