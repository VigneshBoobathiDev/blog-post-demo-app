package controllers_test

import (
	"blogpost/controllers"
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

type CommentControllerTestSuite struct {
	suite.Suite
	mockService       *mocks.MockCommentService
	commentController controllers.CommentController
}

func (suite *CommentControllerTestSuite) SetupTest() {
	suite.mockService = new(mocks.MockCommentService)
	suite.commentController = controllers.NewCommentController(suite.mockService)
}

// Test adding a comment
func (suite *CommentControllerTestSuite) TestAddComment() {
	comment := &models.Comment{
		Nickname: "test_user",
		Comment:  "This is a test comment",
	}

	articleID := 1
	suite.mockService.On("AddComment", articleID, comment).Return(nil)

	// Prepare the request body
	body, _ := json.Marshal(comment)
	req := httptest.NewRequest(http.MethodPost, "/comment/article/1", bytes.NewReader(body))
	req = mux.SetURLVars(req, map[string]string{"id": "1"})

	rec := httptest.NewRecorder()

	// Call the handler
	suite.commentController.AddComment(rec, req)

	// Assert the expectations
	suite.Equal(http.StatusCreated, rec.Code)
	suite.mockService.AssertExpectations(suite.T())
}

// Test adding a reply
func (suite *CommentControllerTestSuite) TestAddReply() {
	replyPayload := map[string]interface{}{
		"parent_comment_id": 1,
		"article_id":        1,
		"reply_comment":     "This is a test reply",
		"nickname":          "test_user",
	}

	suite.mockService.On(
		"AddReply",
		1, 1, "This is a test reply", "test_user",
	).Return(&models.Comment{

		CommentID:       2,
		Nickname:        "test_user",
		Comment:         "This is a test reply",
		ParentCommentID: &[]int{1}[0],
		ArticleID:       1,
	}, nil)

	// Prepare the request body
	body, _ := json.Marshal(replyPayload)
	req := httptest.NewRequest(http.MethodPost, "/comments/reply", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	// Call the handler
	suite.commentController.AddReply(rec, req)

	// Assert the expectations
	suite.Equal(http.StatusOK, rec.Code)
	suite.mockService.AssertExpectations(suite.T())
}

// Test fetching comments by article ID
func (suite *CommentControllerTestSuite) TestGetCommentsByArticleID() {
	articleID := 1
	comments := []models.Comment{
		{CommentID: 1, Nickname: "user1", Comment: "Test comment 1", ArticleID: articleID},
		{CommentID: 2, Nickname: "user2", Comment: "Test comment 2", ArticleID: articleID},
	}

	suite.mockService.On("GetCommentsByArticleID", articleID).Return(comments, nil)

	req := httptest.NewRequest(http.MethodGet, "/comments/1", nil)
	req = mux.SetURLVars(req, map[string]string{"article_id": "1"})

	rec := httptest.NewRecorder()

	// Call the handler
	suite.commentController.GetCommentsByArticleID(rec, req)

	// Assert the expectations
	suite.Equal(http.StatusOK, rec.Code)

	// Check response body
	var response []models.Comment
	json.Unmarshal(rec.Body.Bytes(), &response)
	suite.Equal(comments, response)

	suite.mockService.AssertExpectations(suite.T())
}

// Run the test suite
func TestCommentControllerTestSuite(t *testing.T) {
	suite.Run(t, new(CommentControllerTestSuite))
}
