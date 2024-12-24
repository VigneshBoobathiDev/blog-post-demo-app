package routes

import (
	"blogpost/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(articleController controllers.ArticleController, commentController controllers.CommentController) *mux.Router {
	router := mux.NewRouter()

	// POST /articles
	router.HandleFunc("/create/articles", articleController.CreateArticle).Methods(http.MethodPost)

	// GET /articles/{id}
	router.HandleFunc("/articles/{id}", articleController.GetArticleContent).Methods(http.MethodGet)

	// POST /articles/{id}/comments
	router.HandleFunc("/comment/article/{id}", commentController.AddComment).Methods(http.MethodPost)

	// GET /articles
	router.HandleFunc("/articles", articleController.ListArticles).Methods(http.MethodGet)

	// POST /comments/reply
	router.HandleFunc("/comments/reply", commentController.AddReply).Methods(http.MethodPost)

	// GET /comments/{article_id}
	router.HandleFunc("/comments/{article_id}", commentController.GetCommentsByArticleID).Methods(http.MethodGet)

	return router
}
