package controllers

import (
	"blogpost/models"
	"blogpost/services"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type CommentController interface {
	AddComment(w http.ResponseWriter, r *http.Request)
	AddReply(w http.ResponseWriter, r *http.Request)
	GetCommentsByArticleID(w http.ResponseWriter, r *http.Request)
}

type commentController struct {
	Service services.CommentServiceInterface
}

// Constructor now accepts the interface
func NewCommentController(service services.CommentServiceInterface) CommentController {
	return &commentController{Service: service}
}

// Handler to add a comment to an article
func (cc *commentController) AddComment(w http.ResponseWriter, r *http.Request) {
	// Extract the article ID from the URL
	vars := mux.Vars(r)
	idParam, exists := vars["id"]
	if !exists {
		// Set the response in json format
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Article ID is required"})
		return
	}

	// Convert ID to integer
	articleID, err := strconv.Atoi(idParam)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid Article ID"})
		return
	}

	// Parse the request body to get the comment data
	var comment models.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	// Validate that the "content" field is not empty
	if comment.Comment == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Comment is required"})
		return
	}

	// Add the comment using the service layer
	if err := cc.Service.AddComment(articleID, &comment); err != nil {
		if err.Error() == "article not found" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Article not found"})
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to add comment"})
		}
		return
	}

	// Return the created comment as a response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comment)
}

// Handler to add a reply to a comment
func (cc *commentController) AddReply(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		ParentCommentID int    `json:"parent_comment_id"`
		ArticleID       int    `json:"article_id"`
		ReplyComment    string `json:"reply_comment"`
		Nickname        string `json:"nickname"`
	}

	// Parse the request body
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request payload"})
		return
	}

	// Validate the input
	if payload.ReplyComment == "" || payload.Nickname == "" || payload.ArticleID == 0 || payload.ParentCommentID == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Missing required fields"})
		return
	}

	// Add the reply using the service layer
	reply, err := cc.Service.AddReply(payload.ParentCommentID, payload.ArticleID, payload.ReplyComment, payload.Nickname)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Parent comment not found"})
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "Failed to add reply"})
		}
		return
	}

	// Respond with the created reply
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reply)
}

// Handler to get all comments for an article
func (cc *commentController) GetCommentsByArticleID(w http.ResponseWriter, r *http.Request) {
	// Extract the article ID from the request parameters
	vars := mux.Vars(r)
	articleID, err := strconv.Atoi(vars["article_id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	// Fetch comments using the service
	comments, err := cc.Service.GetCommentsByArticleID(articleID)
	if err != nil {
		//response in json format
		if err.Error() == "article not found" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Article not found"})
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "Failed to fetch comments"})
		return
	}

	// Respond with the comments in JSON format
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
}
