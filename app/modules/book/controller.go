package book

import (
	"gg/domain"
	_ "gg/utils/dto"

	"github.com/gin-gonic/gin"
)

type BookController struct {
	bookSvc domain.IBookService
}

// GetAllBooks godoc
//
// @Summary Get all books
// @Accept json
// @Produce json
// @Success 200 {object} dto.Response[domain.Book]
// @Failure 400 {object} dto.Response[any]
// @Router /book/ [get]
func (bc BookController) GetAllBooks(c *gin.Context) {
	bc.bookSvc.GetAllBooks(c)
}

// SaveBook godoc
//
// @Summary Save book
// @Accept mpfd
// @Produce json
// @Param image formData file true "Image to be uploaded"
// @Param name formData string true "Name of book"
// @Success 200 {object} dto.Response[any]
// @Failude 400 {object} dto.Response[any]
// @Router /book/save [post]
func (bc BookController) SaveBook(c *gin.Context) {
	bc.bookSvc.SaveBook(c)
}

func (bc BookController) GetBook(c *gin.Context) {
	bc.bookSvc.GetBook(c)
}

func (bc BookController) DeleteBookByID(c *gin.Context) {
	bc.bookSvc.DeleteBookByID(c)
}
