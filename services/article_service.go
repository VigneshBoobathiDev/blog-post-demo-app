package services

import (
	"blogpost/logger"
	"blogpost/models"
	"errors"

	"gorm.io/gorm"
)

// Define the interface for the ArticleService
type ArticleServiceInterface interface {
	CreateArticle(article *models.Article) error
	GetArticleByID(articleID int) (*models.Article, error)
	ListArticles(page, pageSize int) ([]models.Article, error)
}

type ArticleService struct {
	DB *gorm.DB
}

// Constructor for the ArticleService
func NewArticleService(db *gorm.DB) *ArticleService {
	return &ArticleService{DB: db}
}

// Create a new article
func (s *ArticleService) CreateArticle(article *models.Article) error {
	logger.Log.Infof("Creating article: %v", article)
	result := s.DB.Create(&article)
	if result.Error != nil {
		logger.Log.Errorf("Failed to create article: %v", result.Error)
		return result.Error
	}
	logger.Log.Infof("Article created successfully: %v", article)
	return nil
}

   // Get an article by ID
   func (s *ArticleService) GetArticleByID(articleID int) (*models.Article, error) {
	logger.Log.Infof("Getting article by ID: %d", articleID)
	var article models.Article

	// Query the database for the article by ID
	result := s.DB.First(&article, articleID)

	if result.Error != nil {
		logger.Log.Errorf("Failed to get article by ID %d: %v", articleID, result.Error)
		return nil, errors.New("article not found")
	}

	return &article, nil
}

// Function to list all articles with pagination
func (s *ArticleService) ListArticles(page, pageSize int) ([]models.Article, error) {
	logger.Log.Infof("Listing articles with page: %d, pageSize: %d", page, pageSize)
	var articles []models.Article

	offset := (page - 1) * pageSize

	// Retrieve articles with pagination
	result := s.DB.Limit(pageSize).Offset(offset).Order("created_at ASC").Find(&articles)
	if result.Error != nil {
		logger.Log.Errorf("Failed to list articles: %v", result.Error)
		return nil, result.Error
	}

	return articles, nil
}
