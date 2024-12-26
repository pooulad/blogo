package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pooulad/blogo/internal/config"
	"github.com/pooulad/blogo/internal/database"
	"github.com/pooulad/blogo/internal/database/model"
	"golang.org/x/crypto/bcrypt"
)

const (
	defaultRole = "user"
)

type App interface {
	// return app config
	GetConfig() *config.Config
	// user crud
	GetAllUsers(ctx *gin.Context) (*[]model.UserResponse, error)
	CreateUser(ctx *gin.Context, user *model.User) error
	UpdateUserByID(ctx *gin.Context, userID int) error
	DeleteUserByID(ctx *gin.Context, userID int) error
	GetUserByID(ctx *gin.Context, userID int) (*model.User, error)
	FollowUserByID(ctx *gin.Context) error
	UnFollowUserByID(ctx *gin.Context) error
	GetFollowersByID(ctx *gin.Context, userID int) (*[]model.UserResponse, error)
	GetFollowingByID(ctx *gin.Context, userID int) (*[]model.UserResponse, error)
	// post crud
	CreatePost(ctx *gin.Context, post *model.Post) error
	GetAllPosts(ctx *gin.Context) (*[]model.PostResponse, error)
	UpdatePostByID(ctx *gin.Context, postID int) error
	DeletePostByID(ctx *gin.Context, postID int) error
	GetPostByID(ctx *gin.Context, postID int) (*model.PostResponse, error)
	LikePostByID(ctx *gin.Context) error
	UnlikePostByID(ctx *gin.Context) error
	// authentication
	Register(ctx *gin.Context, registerInput model.RegisterInput) (*model.UserResponse, error)
	Login(ctx *gin.Context, loginInput model.LoginInput) (*model.UserResponse, error)
}

type app struct {
	store  *database.Store
	config *config.Config
}

func New(store *database.Store, config *config.Config) *app {
	return &app{
		store:  store,
		config: config,
	}
}

func (a *app) GetConfig() *config.Config {
	return a.config
}

func (a *app) GetAllUsers(ctx *gin.Context) (*[]model.UserResponse, error) {
	users, err := a.store.Model.UserResponse.GetAllUsers(a.store.DB)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (a *app) GetFollowersByID(ctx *gin.Context, userID int) (*[]model.UserResponse, error) {
	followers, err := a.store.Model.User.GetFollowers(a.store.DB, userID)
	if err != nil {
		return nil, err
	}

	return followers, nil
}

func (a *app) GetFollowingByID(ctx *gin.Context, userID int) (*[]model.UserResponse, error) {
	following, err := a.store.Model.User.GetFollowings(a.store.DB, userID)
	if err != nil {
		return nil, err
	}

	return following, nil
}

func (a *app) CreateUser(ctx *gin.Context, userBody *model.User) error {
	isUserExist, err := a.store.Model.User.IsUserExistByUsername(a.store.DB, userBody)
	if err != nil {
		return err
	}

	if isUserExist {
		return fmt.Errorf("user already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("internal server error")
	}

	userBody.Password = string(hashedPassword)

	err = a.store.Model.User.CreateUser(a.store.DB, userBody)
	if err != nil {
		return err
	}

	return nil
}

func (a *app) UpdateUserByID(ctx *gin.Context, userID int) error {
	var user model.User
	var updateFields map[string]interface{}

	if err := ctx.ShouldBindJSON(&updateFields); err != nil {
		return fmt.Errorf("invalid request")
	}

	if err := a.store.DB.Where("id=?", userID).Find(&user).Error; err != nil {
		return err
	}

	if user.ID == 0 {
		return fmt.Errorf("user not found")
	}

	if firstName, ok := updateFields["first_name"]; ok {
		user.FirstName = firstName.(string)
	}
	if lastName, ok := updateFields["last_name"]; ok {
		user.LastName = lastName.(string)
	}
	if active, ok := updateFields["active"]; ok {
		user.Active = active.(bool)
	}
	if email, ok := updateFields["email"]; ok {
		user.Email = email.(string)
	}
	if role, ok := updateFields["role"]; ok {
		user.Role = role.(string)
	}
	if skill, ok := updateFields["skill"]; ok {
		user.Skill = skill.(string)
	}
	if password, ok := updateFields["password"]; ok {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password.(string)), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("internal server error")
		}
		user.Password = string(hashedPassword)
	}

	err := a.store.Model.User.UpdateUserByID(a.store.DB, &user)
	if err != nil {
		return err
	}

	return nil
}

func (a *app) GetUserByID(ctx *gin.Context, userID int) (*model.User, error) {
	user, err := a.store.Model.User.GetUserByID(a.store.DB, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *app) FollowUserByID(ctx *gin.Context) error {
	type body struct {
		FollowerID uint `json:"follower_id"`
		FollowedID uint `json:"followed_id"`
	}

	var requestBody body

	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		return err
	}

	err = a.store.Model.User.FollowUserByID(a.store.DB, requestBody.FollowerID, requestBody.FollowedID)
	if err != nil {
		return err
	}

	return nil
}

func (a *app) UnFollowUserByID(ctx *gin.Context) error {
	type body struct {
		FollowerID uint `json:"follower_id"`
		FollowedID uint `json:"followed_id"`
	}

	var requestBody body

	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		return err
	}

	err = a.store.Model.User.UnFollowUserByID(a.store.DB, requestBody.FollowerID, requestBody.FollowedID)
	if err != nil {
		return err
	}

	return nil
}

func (a *app) DeleteUserByID(ctx *gin.Context, userID int) error {
	err := a.store.Model.User.DeleteUserByID(a.store.DB, userID)
	if err != nil {
		return err
	}

	return nil
}

func (a *app) CreatePost(ctx *gin.Context, postBody *model.Post) error {
	err := a.store.Model.Post.CreatePost(a.store.DB, postBody)
	if err != nil {
		return err
	}

	return nil
}

func (a *app) GetAllPosts(ctx *gin.Context) (*[]model.PostResponse, error) {
	var response []model.PostResponse
	posts, err := a.store.Model.Post.GetAllPosts(a.store.DB)
	if err != nil {
		return nil, err
	}

	username, exist := ctx.Get("username")
	if username == nil {
		return nil, fmt.Errorf("user id not found in context")
	}

	var user model.User
	user.Username = username.(string)
	isUserExist, err := a.store.Model.User.IsUserExistByUsername(a.store.DB, &user)
	if err != nil {
		return nil, fmt.Errorf("get current user data faild")
	}

	if !isUserExist {
		return nil, fmt.Errorf("get current user data faild")
	}

	for _, post := range *posts {
		liked := false
		for _, userRefer := range (post).LikedBy {
			if exist {
				if userRefer.ID == user.ID {
					liked = true
					break
				}
			}
		}

		response = append(response, model.PostResponse{
			ID:         post.ID,
			Title:      post.Title,
			Content:    post.Content,
			UserRefer:  post.UserRefer,
			Liked:      liked,
			LikedCount: len(post.LikedBy),
		})
	}

	return &response, nil
}

func (a *app) UpdatePostByID(ctx *gin.Context, postID int) error {
	var post model.Post
	var updateFields map[string]interface{}

	if err := ctx.ShouldBindJSON(&updateFields); err != nil {
		return fmt.Errorf("invalid request")
	}

	if err := a.store.DB.Where("id=?", postID).Find(&post).Error; err != nil {
		return err
	}

	if post.ID == 0 {
		return fmt.Errorf("post not found")
	}

	if title, ok := updateFields["title"]; ok {
		post.Title = title.(string)
	}

	if content, ok := updateFields["content"]; ok {
		post.Content = content.(string)
	}

	if userRefer, ok := updateFields["user_id"]; ok {
		post.UserRefer = userRefer.(uint)
	}

	err := a.store.Model.Post.UpdatePostByID(a.store.DB, &post)
	if err != nil {
		return err
	}

	return nil
}

func (a *app) DeletePostByID(ctx *gin.Context, postID int) error {
	err := a.store.Model.Post.DeletePostByID(a.store.DB, postID)
	if err != nil {
		return err
	}

	return nil
}

func (a *app) GetPostByID(ctx *gin.Context, postID int) (*model.PostResponse, error) {
	var response model.PostResponse
	username, exist := ctx.Get("username")
	if username == nil {
		return nil, fmt.Errorf("user id not found in context")
	}

	var user model.User
	user.Username = username.(string)
	isUserExist, err := a.store.Model.User.IsUserExistByUsername(a.store.DB, &user)
	if err != nil {
		return nil, fmt.Errorf("get current user data faild")
	}

	if !isUserExist {
		return nil, fmt.Errorf("get current user data faild")
	}

	post, err := a.store.Model.Post.GetPostByID(a.store.DB, postID)
	if err != nil {
		return nil, err
	}

	liked := false
	for _, userRefer := range post.LikedBy {
		if exist {
			if userRefer.ID == user.ID {
				liked = true
				break
			}
		}
	}

	response.ID = post.ID
	response.Title = post.Title
	response.Content = post.Content
	response.UserRefer = post.UserRefer
	response.Liked = liked
	response.LikedCount = len(post.LikedBy)

	return &response, nil
}

func (a *app) LikePostByID(ctx *gin.Context) error {
	type body struct {
		UserID uint `json:"user_id"`
		PostID uint `json:"post_id"`
	}

	var requestBody body

	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		return err
	}

	err = a.store.Model.Post.LikePostByID(a.store.DB, requestBody.UserID, requestBody.PostID)
	if err != nil {
		return err
	}

	return nil
}

func (a *app) UnlikePostByID(ctx *gin.Context) error {
	type body struct {
		UserID uint `json:"user_id"`
		PostID uint `json:"post_id"`
	}

	var requestBody body

	err := ctx.ShouldBindJSON(&requestBody)
	if err != nil {
		return err
	}

	err = a.store.Model.Post.UnlikePostByID(a.store.DB, requestBody.UserID, requestBody.PostID)
	if err != nil {
		return err
	}

	return nil
}

func (a *app) Register(ctx *gin.Context, registerInput model.RegisterInput) (*model.UserResponse, error) {
	var user model.User
	var userResponse model.UserResponse
	if err := a.store.DB.Where("username=?", registerInput.Username).Find(&user).Error; err != nil {
		return nil, err
	}

	if user.ID != 0 {
		return nil, fmt.Errorf("user already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	// fill and create with form data here
	user.FirstName = registerInput.FirstName
	user.LastName = registerInput.LastName
	user.Username = registerInput.Username
	user.Password = string(hashedPassword)
	user.Email = registerInput.Email
	user.Skill = registerInput.Skill
	user.Role = defaultRole

	if err := a.store.DB.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("internal server error")
	}

	// fill user response data here
	userResponse.ID = user.ID
	userResponse.FirstName = registerInput.FirstName
	userResponse.LastName = registerInput.LastName
	userResponse.Username = registerInput.Username
	userResponse.Email = registerInput.Email
	userResponse.Skill = registerInput.Skill
	userResponse.Role = defaultRole

	return &userResponse, nil
}

func (a *app) Login(ctx *gin.Context, loginInput model.LoginInput) (*model.UserResponse, error) {
	var user model.User
	var userResponse model.UserResponse
	if err := a.store.DB.Where("username=?", loginInput.Username).Find(&user).Error; err != nil {
		return nil, fmt.Errorf("database error")
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	compaireErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password))
	if compaireErr != nil {
		return nil, fmt.Errorf("password is invalid")
	}

	userResponse.ID = user.ID
	userResponse.FirstName = user.FirstName
	userResponse.LastName = user.LastName
	userResponse.Username = user.Username
	userResponse.Email = user.Email
	userResponse.Skill = user.Skill
	userResponse.Email = user.Email
	userResponse.LastVisited = user.LastVisited

	return &userResponse, nil
}
