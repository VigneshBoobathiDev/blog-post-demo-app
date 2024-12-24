package services

import (
	"blogpost/logger"
	"blogpost/models"
	"errors"

	"gorm.io/gorm"
)

// Define the interface for the CommentService
type CommentServiceInterface interface {
	AddComment(articleID int, comment *models.Comment) error
	AddReply(parentCommentID, articleID int, replyComment, nickname string) (*models.Comment, error)
	GetCommentsByArticleID(articleID int) ([]models.Comment, error)
}

type CommentService struct {
	DB *gorm.DB
}

// Constructor for CommentService
func NewCommentService(db *gorm.DB) *CommentService {
	return &CommentService{DB: db}
}

func (c *CommentService) AddComment(articleID int, comment *models.Comment) error {
	logger.Log.Infof("Adding comment to article ID: %d", articleID)
	var article models.Article
	if err := c.DB.First(&article, articleID).Error; err != nil {
		logger.Log.Errorf("Failed to find article ID %d: %v", articleID, err)
		return errors.New("article not found")
	}
	comment.ArticleID = articleID
	if err := c.DB.Create(comment).Error; err != nil {
		logger.Log.Errorf("Failed to add comment to article ID %d: %v", articleID, err)
		return err
	}

	logger.Log.Infof("Comment successfully added to article ID: %d", articleID)
	return nil
}

func (c *CommentService) AddReply(parentCommentID, articleID int, replyComment, nickname string) (*models.Comment, error) {
	logger.Log.Infof("Adding reply to parent comment ID: %d", parentCommentID)
	var parentComment models.Comment
	if err := c.DB.Where("comment_id = ? AND article_id = ?", parentCommentID, articleID).First(&parentComment).Error; err != nil {
		logger.Log.Errorf("Failed to find parent comment ID %d: %v", parentCommentID, err)
		return nil, errors.New("parent comment not found")
	}
	reply := models.Comment{
		ArticleID:       articleID,
		ReplyComment:    &replyComment,
		Nickname:        nickname,
		ParentCommentID: &parentCommentID,
	}
	if err := c.DB.Create(&reply).Error; err != nil {
		logger.Log.Errorf("Failed to add reply to parent comment ID %d: %v", parentCommentID, err)
		return nil, err
	}

	logger.Log.Infof("Reply successfully added to parent comment ID: %d", parentCommentID)
	return &reply, nil
}

func (c *CommentService) GetCommentsByArticleID(articleID int) ([]models.Comment, error) {
	logger.Log.Infof("Getting comments for article ID: %d", articleID)	
	var comments []models.Comment
	if err := c.DB.Preload("Replies").Where("article_id = ? AND parent_comment_id IS NULL", articleID).Find(&comments).Error; err != nil {
		logger.Log.Errorf("Failed to get comments for article ID %d: %v", articleID, err)
		return nil, err
	}
	return comments, nil
}
