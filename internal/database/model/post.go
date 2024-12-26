package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserRefer uint   `json:"user_id"`
	LikedBy   []User `gorm:"many2many:likes;"`
}

type PostResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	UserRefer  uint   `json:"user_id"`
	Liked      bool   `json:"liked"`
	LikedCount int    `json:"liked_count"`
}

func (p *Post) CreatePost(db *gorm.DB, postBody *Post) error {
	post := Post{
		Title:     postBody.Title,
		Content:   postBody.Content,
		UserRefer: postBody.UserRefer,
	}

	if err := db.Create(&post).Error; err != nil {
		return err
	}
	return nil
}

func (p *Post) GetAllPosts(db *gorm.DB) (*[]Post, error) {
	var posts *[]Post
	if err := db.Table("posts").Preload("LikedBy").Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *Post) UpdatePostByID(db *gorm.DB, postBody *Post) error {
	if err := db.Save(&postBody).Error; err != nil {
		return err
	}

	return nil
}

func (p *Post) DeletePostByID(db *gorm.DB, postID int) error {
	var post *Post
	if err := db.Table("posts").Where("id=?", postID).Find(&post).Error; err != nil {
		return err
	}

	if post.ID == 0 {
		return fmt.Errorf("post not found")
	}

	if err := db.Table("posts").Delete(&post).Error; err != nil {
		return err
	}

	return nil
}

func (p *Post) GetPostByID(db *gorm.DB, postID int) (*Post, error) {
	var post *Post
	if err := db.Table("posts").Where("id=?", postID).Preload("LikedBy").Find(&post).Error; err != nil {
		return nil, err
	}

	if post.ID == 0 {
		return nil, fmt.Errorf("post not found")
	}

	return post, nil
}

func (p *Post) LikePostByID(db *gorm.DB, userID, postID uint) error {
	return db.Table("likes").Create(map[string]interface{}{
		"user_id": userID,
		"post_id": postID,
	}).Error
}

func (p *Post) UnlikePostByID(db *gorm.DB, userID, postID uint) error {
	return db.Table("likes").Where("user_id = ? AND post_id = ?", userID, postID).Delete(nil).Error
}
