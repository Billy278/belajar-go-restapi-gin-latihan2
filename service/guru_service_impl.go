package service

import (
	"belajar-go-restapi-gin-latihan2/model/domain"
	"belajar-go-restapi-gin-latihan2/model/web"
	"belajar-go-restapi-gin-latihan2/repository"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GuruServiceImpl struct {
	DB         *sql.DB
	Repository repository.GuruRepository
}

func NewGuruServiceImpl(db *sql.DB, repostitory repository.GuruRepository) GuruService {
	return &GuruServiceImpl{
		DB:         db,
		Repository: repostitory,
	}
}
func (guru_service *GuruServiceImpl) Create(c *gin.Context) {
	createGuru := web.CreateGuru{}
	db := guru_service.DB
	err := c.ShouldBindJSON(&createGuru)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponses{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		})
		return
	}
	guru := domain.Guru{
		Id_guru: createGuru.Id_guru,
		Name:    createGuru.Name,
		Status:  createGuru.Status,
	}
	guru = guru_service.Repository.Create(c, db, guru)
	c.JSON(http.StatusOK, web.WebResponses{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   guru,
	})

}

func (guru_service *GuruServiceImpl) Update(c *gin.Context) {
	updateGuru := web.UpdateGuru{}
	id := c.Param("IdGuru")
	db := guru_service.DB
	err := c.ShouldBindJSON(&updateGuru)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, web.WebResponses{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err,
		})
		return
	}
	guru, err := guru_service.Repository.FindById(c, db, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, web.WebResponses{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   err,
		})
		return
	}
	guru.Id_guru = id
	guru.Name = updateGuru.Name
	guru.Status = updateGuru.Status

	guru = guru_service.Repository.Update(c, db, guru)
	c.JSON(http.StatusOK, web.WebResponses{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   guru,
	})
}

func (guru_service *GuruServiceImpl) Delete(c *gin.Context) {
	db := guru_service.DB
	id := c.Param("IdGuru")
	_, err := guru_service.Repository.FindById(c, db, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, web.WebResponses{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   err,
		})
		return
	}
	guru_service.Repository.Delete(c, db, id)
	c.JSON(http.StatusOK, web.WebResponses{
		Code:   http.StatusOK,
		Status: "Ok",
	})

}

func (guru_service *GuruServiceImpl) FindById(c *gin.Context) {
	db := guru_service.DB
	id := c.Param("IdGuru")
	guru, err := guru_service.Repository.FindById(c, db, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, web.WebResponses{
			Code:   http.StatusNotFound,
			Status: "Not Found",
			Data:   err,
		})
		return
	}
	c.JSON(http.StatusOK, web.WebResponses{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   guru,
	})
}

func (guru_service *GuruServiceImpl) FindAll(c *gin.Context) {
	db := guru_service.DB
	guru := guru_service.Repository.FindAll(c, db)
	c.JSON(http.StatusOK, web.WebResponses{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   guru,
	})
}
