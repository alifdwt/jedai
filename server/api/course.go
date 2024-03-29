package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	db "github.com/alifdwt/jedai/server/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createCourseRequest struct {
	ID          string `json:"id" binding:"required"`
	UserID      string `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageUrl    string `json:"image_url" binding:"required"`
	Price       int64  `json:"price" binding:"required"`
	IsPublished bool   `json:"is_published"`
	CategoryID  string `json:"category_id" binding:"required"`
}

type courseResponse struct {
	ID          string           `json:"id"`
	User        userResponse     `json:"user"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	ImageUrl    string           `json:"image_url"`
	Price       int64            `json:"price"`
	IsPublished bool             `json:"is_published"`
	Category    categoryResponse `json:"category"`
	CreatedAt   time.Time        `json:"created_at"`
}

func (server *Server) createCourse(ctx *gin.Context) {
	var req createCourseRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCourseParams{
		ID:          req.ID,
		UserID:      req.UserID,
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
		Price:       req.Price,
		IsPublished: req.IsPublished,
		CategoryID:  req.CategoryID,
	}

	course, err := server.store.CreateCourse(ctx, arg)
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

	ctx.JSON(http.StatusOK, course)
}

type getCourseRequest struct {
	ID     string `uri:"id" binding:"required"`
	UserID string `uri:"user_id" binding:"required"`
}

func (server *Server) getCourse(ctx *gin.Context) {
	var req getCourseRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.GetCourseParams{
		ID:     req.ID,
		UserID: req.UserID,
	}

	course, err := server.store.GetCourse(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := courseResponse{
		ID: course.ID,
		User: userResponse{
			Username:  course.UserID,
			FullName:  course.FullName,
			Email:     course.Email,
			ImageUrl:  course.ImageUrl_2,
			BannerUrl: course.BannerUrl,
		},
		Title:       course.Title,
		Description: course.Description,
		ImageUrl:    course.ImageUrl,
		Price:       course.Price,
		IsPublished: course.IsPublished,
		Category: categoryResponse{
			ID:   course.CategoryID,
			Name: course.Name,
		},
		CreatedAt: course.CreatedAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}

type listCoursesRequest struct {
	PageID     int32  `form:"page_id" binding:"required,min=1"`
	PageSize   int32  `form:"page_size" binding:"required,min=5"`
	CategoryID string `form:"category_id"`
}

func (server *Server) listCourses(ctx *gin.Context) {
	var req listCoursesRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCoursesParams{
		Limit:      req.PageSize,
		Offset:     (req.PageID - 1) * req.PageSize,
		CategoryID: fmt.Sprintf("%%%s%%", req.CategoryID),
	}

	courses, err := server.store.ListCourses(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var listCoursesResponse []courseResponse
	for _, course := range courses {
		listCoursesResponse = append(listCoursesResponse, courseResponse{
			ID: course.ID,
			User: userResponse{
				Username: course.UserID,
				FullName: course.FullName,
				Email:    course.Email,
				ImageUrl: course.ImageUrl_2,
			},
			Title:       course.Title,
			Description: course.Description,
			ImageUrl:    course.ImageUrl,
			Price:       course.Price,
			IsPublished: course.IsPublished,
			Category: categoryResponse{
				ID:   course.CategoryID,
				Name: course.Name,
			},
			CreatedAt: course.CreatedAt,
		})
	}

	ctx.JSON(http.StatusOK, listCoursesResponse)
}
