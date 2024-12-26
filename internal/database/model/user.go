package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string    `json:"first_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	Role        string    `json:"role"`
	Active      bool      `json:"active"`
	Skill       string    `json:"skill"`
	LastVisited time.Time `json:"last_visited,omitempty"`
	Posts       []Post    `json:"posts" gorm:"foreignKey:UserRefer"`
	LikedPosts  []Post    `gorm:"many2many:likes;"`
	Followers   []User    `gorm:"many2many:user_follows;joinForeignKey:FollowedID;joinReferences:FollowerID" json:"followers,omitempty"`
	Following   []User    `gorm:"many2many:user_follows;joinForeignKey:FollowerID;joinReferences:FollowedID" json:"following,omitempty"`
}

type UserResponse struct {
	gorm.Model
	FirstName   string         `json:"first_name,omitempty"`
	LastName    string         `json:"last_name,omitempty"`
	Username    string         `json:"username"`
	Email       string         `json:"email"`
	Role        string         `json:"role"`
	Skill       string         `json:"skill"`
	LastVisited time.Time      `json:"last_visited,omitempty"`
	Posts       []Post         `json:"posts" gorm:"foreignKey:UserRefer"`
	Followers   []UserResponse `gorm:"many2many:user_follows;joinForeignKey:FollowedID;joinReferences:FollowerID" json:"followers,omitempty"`
	Following   []UserResponse `gorm:"many2many:user_follows;joinForeignKey:FollowerID;joinReferences:FollowedID" json:"following,omitempty"`
}

func (u *User) CreateUser(db *gorm.DB, userBody *User) error {
	if err := db.Create(&userBody).Error; err != nil {
		return err
	}

	return nil
}

func (u *User) UpdateUserByID(db *gorm.DB, userBody *User) error {
	if err := db.Save(&userBody).Error; err != nil {
		return err
	}

	return nil
}

func (u *User) DeleteUserByID(db *gorm.DB, userID int) error {
	var user *User
	if err := db.Table("users").Where("id=?", userID).Find(&user).Error; err != nil {
		return err
	}

	if user.ID == 0 {
		return fmt.Errorf("user not found")
	}

	if err := db.Table("users").Select("Posts").Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

func (u *User) GetUserByID(db *gorm.DB, userID int) (*User, error) {
	var user *User
	if err := db.Table("users").Where("id=?", userID).Preload("Posts").Find(&user).Error; err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (u *UserResponse) GetAllUsers(db *gorm.DB) (*[]UserResponse, error) {
	var users *[]UserResponse
	if err := db.Table("users").Select("id", "first_name", "last_name", "username", "email", "role", "skill").Preload("Posts").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (u *User) IsUserExistByUsername(db *gorm.DB, userBody *User) (bool, error) {
	if err := db.Where("username=?", &userBody.Username).Find(&userBody).Error; err != nil {
		return false, err
	}

	if userBody.ID != 0 {
		return true, nil
	}

	return false, nil
}

func (u *User) FollowUserByID(db *gorm.DB, followerID, followedID uint) error {
	return db.Table("user_follows").Create(map[string]interface{}{
		"follower_id": followerID,
		"followed_id": followedID,
	}).Error
}

func (u *User) UnFollowUserByID(db *gorm.DB, followerID, followedID uint) error {
	return db.Table("user_follows").Where("follower_id = ? AND followed_id = ?", followerID, followedID).Delete(nil).Error
}

func (s *User) GetFollowers(db *gorm.DB, userID int) (*[]UserResponse, error) {
	var user User
	var response []UserResponse

	if err := db.Table("users").Preload("Followers.Posts").First(&user, userID).Error; err != nil {
		return nil, err
	}

	for _, user := range user.Followers {
		var userResponse UserResponse

		userResponse.ID = user.ID
		userResponse.FirstName = user.FirstName
		userResponse.LastName = user.LastName
		userResponse.Username = user.Username
		userResponse.Email = user.Email
		userResponse.Skill = user.Skill
		userResponse.Role = user.Role
		userResponse.Posts = user.Posts

		response = append(response, userResponse)
	}

	return &response, nil
}

func (u *User) GetFollowings(db *gorm.DB, userID int) (*[]UserResponse, error) {
	var user User
	var response []UserResponse

	if err := db.Table("users").Preload("Following.Posts").First(&user, userID).Error; err != nil {
		return nil, err
	}

	for _, user := range user.Following {
		var userResponse UserResponse

		userResponse.ID = user.ID
		userResponse.FirstName = user.FirstName
		userResponse.LastName = user.LastName
		userResponse.Username = user.Username
		userResponse.Email = user.Email
		userResponse.Skill = user.Skill
		userResponse.Role = user.Role
		userResponse.Posts = user.Posts

		response = append(response, userResponse)
	}

	return &response, nil
}
