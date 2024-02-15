package api

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	db "github.com/alifdwt/jedai/server/db/sqlc"
	"github.com/alifdwt/jedai/server/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type userResponse struct {
	Username          string    `json:"username"`
	FullName          string    `json:"full_name"`
	Email             string    `json:"email"`
	ImageUrl          string    `json:"image_url"`
	BannerUrl         string    `json:"banner_url"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type createUserRequest struct {
	Username  string `json:"username" binding:"required,alphanum"`
	Password  string `json:"password" binding:"required,min=6"`
	FullName  string `json:"full_name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	ImageUrl  string `json:"image_url"`
	BannerUrl string `json:"banner_url"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Username:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		ImageUrl:          user.ImageUrl,
		BannerUrl:         user.BannerUrl,
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

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
		ImageUrl:       req.ImageUrl,
		BannerUrl:      req.BannerUrl,
	}

	user, err := server.store.CreateUser(ctx, arg)
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

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type getUserRequest struct {
	Username string `uri:"username" binding:"required,alphanum"`
}

type getUserResponse struct {
	Username  string      `json:"username"`
	FullName  string      `json:"full_name"`
	Email     string      `json:"email"`
	ImageUrl  string      `json:"image_url"`
	BannerUrl string      `json:"banner_url"`
	CreatedAt time.Time   `json:"created_at"`
	Courses   []db.Course `json:"courses"`
}

func (server *Server) getUser(ctx *gin.Context) {
	var req getUserRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUserWithCourses(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var jsonData = []byte(user.Courses)
	var courses []db.Course

	err = json.Unmarshal(jsonData, &courses)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := getUserResponse{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		ImageUrl:  user.ImageUrl,
		BannerUrl: user.BannerUrl,
		CreatedAt: user.CreatedAt,
		Courses:   courses,
	}

	ctx.JSON(http.StatusOK, rsp)
}

type listUsersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5"`
}

func (server *Server) listUsers(ctx *gin.Context) {
	var req listUsersRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var listUsersResponse []userResponse
	for _, user := range users {
		listUsersResponse = append(listUsersResponse, userResponse{
			Username:          user.Username,
			FullName:          user.FullName,
			Email:             user.Email,
			ImageUrl:          user.ImageUrl,
			BannerUrl:         user.BannerUrl,
			PasswordChangedAt: user.PasswordChangedAt,
			CreatedAt:         user.CreatedAt,
		})
	}

	ctx.JSON(http.StatusOK, listUsersResponse)
}
