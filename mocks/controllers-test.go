package mocks

import (
	"blogpost/models"

	"github.com/stretchr/testify/mock"
)

type MockArticleService struct {
	mock.Mock
}

type MockCommentService struct {
	mock.Mock
}

func (m *MockArticleService) CreateArticle(article *models.Article) error {
	args := m.Called(article)
	return args.Error(0)
}

func (m *MockArticleService) GetArticleByID(articleID int) (*models.Article, error) {
	args := m.Called(articleID)
	return args.Get(0).(*models.Article), args.Error(1)
}

func (m *MockArticleService) ListArticles(page, pageSize int) ([]models.Article, error) {
	args := m.Called(page, pageSize)
	return args.Get(0).([]models.Article), args.Error(1)
}


func (m *MockCommentService) AddComment(articleID int, comment *models.Comment) error {
	args := m.Called(articleID, comment)
	return args.Error(0)
}

func (m *MockCommentService) AddReply(parentCommentID, articleID int, replyComment, nickname string) (*models.Comment, error) {
	args := m.Called(parentCommentID, articleID, replyComment, nickname)
	return args.Get(0).(*models.Comment), args.Error(1)
}

func (m *MockCommentService) GetCommentsByArticleID(articleID int) ([]models.Comment, error) {
	args := m.Called(articleID)
	return args.Get(0).([]models.Comment), args.Error(1)
}