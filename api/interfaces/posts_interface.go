package interfaces

import "github.com/victorsteven/fullstack/api/models"

type PostInterface interface {
	SavePost(models.Post) (models.Post, error)
	FindAllPosts() ([]models.Post, error)
	FindPostByID(uint64) (models.Post, error)
	UpdateAPost(uint64, models.Post) (int64, error)
	DeleteAPost(postID uint64, userID uint32) (int64, error)
	// Delete(postID uint64) (int64, error)

}
