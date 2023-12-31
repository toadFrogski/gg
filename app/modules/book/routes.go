package book

// DELETE THIS
// Reason: Separation routes and modules
// Need to incapsulate module behaviour
// "Chain of responsability"

import (
	"gg/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooksRoutes(r gin.IRouter, db *gorm.DB) gin.IRouter {

	bookAPI := Wire(db)

	bookRouter := r.Group("book")
	{
		bookRouter.GET("/", bookAPI.GetAllBooks)
		bookRouter.GET(":bookID", bookAPI.GetBook)
		bookRouter.POST("save", middlewares.JwtAuthMiddleware(), bookAPI.SaveBook)
		bookRouter.DELETE(":bookID", bookAPI.DeleteBookByID)
	}

	return r
}
