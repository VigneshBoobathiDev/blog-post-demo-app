    package models

    import "time"

    type Article struct {
        ArticleID int       `gorm:"column:article_id;primaryKey;autoIncrement" json:"articleId"`
        Nickname  string    `gorm:"size:255;not null" json:"nickname"`
        Title     string    `gorm:"size:255;not null" json:"title"`
        Content   string    `gorm:"type:text;not null" json:"content"`
        CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
    }

    type Comment struct {
        CommentID       int        `gorm:"primaryKey;autoIncrement" json:"comment_id"` // Primary Key
        ArticleID       int        `gorm:"column:article_id;foreignkey:ArticleID" json:"articleId"`
        Comment         string     `gorm:"type:text;not null" json:"comment,omitempty"`
        Nickname        string     `gorm:"size:255;not null" json:"nickname,omitempty"`
        CreatedAt       time.Time  `gorm:"autoCreateTime" json:"created_at"`
        ParentCommentID *int       `json:"parent_comment_id,omitempty"` // Foreign Key to Parent Comment (nullable)
        ReplyComment    *string    `gorm:"type:text" json:"reply_comment,omitempty"`
        Replies         []Comment  `gorm:"foreignKey:ParentCommentID;constraint:OnDelete:CASCADE;" json:"replies,omitempty"`
    }
    