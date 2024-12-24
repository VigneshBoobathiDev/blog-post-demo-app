package controllers

import (
	"blogpost/mocks"
	"blogpost/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gorilla/mux"

	"github.com/stretchr/testify/suite"
)

type ArticleControllerTestSuite struct {
	suite.Suite
	mockService        *mocks.MockArticleService
	articleController  ArticleController
}

func (suite *ArticleControllerTestSuite) SetupTest() {
	suite.mockService = new(mocks.MockArticleService)
	suite.articleController = NewArticleController(suite.mockService)
}

func (suite *ArticleControllerTestSuite) TestCreateArticle() {
	article := models.Article{Nickname: "vicky", Title: "Test Article", Content: "This is a test"}
	suite.mockService.On("CreateArticle", &article).Return(nil)

	body, _ := json.Marshal(article)
	req := httptest.NewRequest(http.MethodPost, "/create/articles", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	suite.articleController.CreateArticle(rec, req)

	suite.Equal(http.StatusCreated, rec.Code)
	suite.mockService.AssertExpectations(suite.T())
}

func (suite *ArticleControllerTestSuite) TestGetArticleContent() {
	article := &models.Article{ArticleID: 1, Content: "Sample Content"}
	suite.mockService.On("GetArticleByID", 1).Return(article, nil)

	// Set up the router
	router := mux.NewRouter()
	router.HandleFunc("/articles/{id}", suite.articleController.GetArticleContent)

	req := httptest.NewRequest(http.MethodGet, "/articles/1", nil)
	rec := httptest.NewRecorder()

	// Serve the request
	router.ServeHTTP(rec, req)

	suite.Equal(http.StatusOK, rec.Code)
	suite.mockService.AssertExpectations(suite.T())
}

func TestArticleControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ArticleControllerTestSuite))
}