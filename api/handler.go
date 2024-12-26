package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pooulad/blogo/internal/database/model"
)

// GetAllUsers godoc
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Success response containing users"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/users [get]
func (a *api) GetAllUsers(ctx *gin.Context) {
	users, err := a.app.GetAllUsers(ctx)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "get users failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"users": users,
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "get users failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User details"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 400 {object} map[string]interface{} "Bad request or validation error"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/users [post]
func (a *api) CreateUser(ctx *gin.Context) {
	var user model.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Create user failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	fmt.Println(user.Password)

	err := a.app.CreateUser(ctx, &user)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Create user failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"success": true,
			"message": "Create user successful",
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Create user failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// DeleteUserByID godoc
// @Summary Delete a user by ID
// @Description Remove a user from the system by their unique ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 400 {object} map[string]interface{} "Bad request or invalid user ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/users/{id} [delete]
func (a *api) DeleteUserByID(ctx *gin.Context) {
	userID, err := GetParamByName(ctx, "id")
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Delete user failed",
				"error":   fmt.Errorf("user id param is invalid").Error(),
			},
		}, fmt.Errorf("user id param is invalid"))
		return
	}

	err = a.app.DeleteUserByID(ctx, userID.(int))
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Delete user failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"success": true,
			"message": "Delete user successful",
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Delete user failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Retrieve a user's details using their unique ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{} "Success response with user details"
// @Failure 400 {object} map[string]interface{} "Bad request or invalid user ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/users/{id} [get]
func (a *api) GetUserByID(ctx *gin.Context) {
	userID, err := GetParamByName(ctx, "id")
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Get user by id failed",
				"error":   fmt.Errorf("user id param is invalid").Error(),
			},
		}, fmt.Errorf("user id param is invalid"))
		return
	}

	user, err := a.app.GetUserByID(ctx, userID.(int))
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Get user by id failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"success": true,
			"message": "Get user by id successful",
			"user":    user,
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Get user by id failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// UpdateUserByID godoc
// @Summary Update a user by ID
// @Description Modify a user's details using their unique ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body map[string]interface{} true "Updated user data"
// @Success 200 {object} map[string]interface{} "Success response"
// @Failure 400 {object} map[string]interface{} "Bad request or invalid user ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/users/{id} [patch]
func (a *api) UpdateUserByID(ctx *gin.Context) {
	userID, err := GetParamByName(ctx, "id")
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Update user failed",
				"error":   fmt.Errorf("user id param is invalid").Error(),
			},
		}, fmt.Errorf("user id param is invalid"))
		return
	}

	err = a.app.UpdateUserByID(ctx, userID.(int))
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Update user failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"success": true,
			"message": "Update user successful",
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Update user failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// GetFollowersByID godoc
// @Summary Get followers by user ID
// @Description Retrieve a list of followers for a specific user using their unique ID
// @Tags followers
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{} "Success response with followers list"
// @Failure 400 {object} map[string]interface{} "Bad request or invalid user ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/users/{id}/followers [get]
func (a *api) GetFollowersByID(ctx *gin.Context) {
	userID, err := GetParamByName(ctx, "id")
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Get followers by id failed",
				"error":   fmt.Errorf("user id param is invalid").Error(),
			},
		}, fmt.Errorf("user id param is invalid"))
		return
	}

	followers, err := a.app.GetFollowersByID(ctx, userID.(int))
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Get followers by id failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"success":   true,
			"message":   "Get followers by id successful",
			"followers": followers,
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Get followers by id failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// GetFollowingByID godoc
// @Summary Get following by user ID
// @Description Retrieve a list of users that a specific user is following using their unique ID
// @Tags following
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{} "Success response with following list"
// @Failure 400 {object} map[string]interface{} "Bad request or invalid user ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/users/{id}/following [get]
func (a *api) GetFollowingByID(ctx *gin.Context) {
	userID, err := GetParamByName(ctx, "id")
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Get following by id failed",
				"error":   fmt.Errorf("user id param is invalid").Error(),
			},
		}, fmt.Errorf("user id param is invalid"))
		return
	}

	following, err := a.app.GetFollowingByID(ctx, userID.(int))
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Get following by id failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"success":   true,
			"message":   "Get following by id successful",
			"following": following,
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Get following by id failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// FollowUserByID godoc
// @Summary Follow a user by their ID
// @Description Allows a user to follow another user by providing the target user's unique ID
// @Tags following
// @Accept json
// @Produce json
// @Param id path int true "Target User ID"
// @Success 200 {object} map[string]interface{} "Success response when follow action is successful"
// @Failure 400 {object} map[string]interface{} "Bad request or invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/users/{id}/follow [post]
func (a *api) FollowUserByID(ctx *gin.Context) {
	err := a.app.FollowUserByID(ctx)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Follow user failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"success": true,
			"message": "Follow user successful",
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Follow user failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// UnFollowUserByID godoc
// @Summary Unfollow a user by their ID
// @Description Allows a user to unfollow another user by providing the target user's unique ID
// @Tags following
// @Accept json
// @Produce json
// @Param id path int true "Target User ID"
// @Success 200 {object} map[string]interface{} "Success response when unfollow action is successful"
// @Failure 400 {object} map[string]interface{} "Bad request or invalid input"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/users/{id}/unfollow [post]
func (a *api) UnFollowUserByID(ctx *gin.Context) {
	err := a.app.UnFollowUserByID(ctx)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Unfollow user failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"success": true,
			"message": "Unfollow user successful",
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Unfollow user failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// GetAllPosts godoc
// @Summary Get all posts
// @Description Retrieve all posts from the application
// @Tags posts
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Success response with a list of posts"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/posts [get]
func (a *api) GetAllPosts(ctx *gin.Context) {
	posts, err := a.app.GetAllPosts(ctx)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "get posts failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"posts": posts,
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "get posts failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// CreatePost godoc
// @Summary Create a new post
// @Description Create a new post with the given data
// @Tags posts
// @Accept json
// @Produce json
// @Param post body model.Post true "Post data"
// @Success 200 {object} map[string]interface{} "Success message for post creation"
// @Failure 400 {object} map[string]interface{} "Bad request error with message"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/posts [post]
func (a *api) CreatePost(ctx *gin.Context) {
	var post *model.Post

	if err := ctx.ShouldBindJSON(&post); err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "missing post form data",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err := a.app.CreatePost(ctx, post)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Create new post failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"success": true,
			"message": "Create new post successful",
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Create new post failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// GetPostByID godoc
// @Summary Get post by ID
// @Description Retrieve a post by its unique ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} map[string]interface{} "Post data"
// @Failure 400 {object} map[string]interface{} "Bad request error with message"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/posts/{id} [get]
func (a *api) GetPostByID(ctx *gin.Context) {
	postID, err := GetParamByName(ctx, "id")
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Get post by id failed",
				"error":   fmt.Errorf("user id param is invalid").Error(),
			},
		}, fmt.Errorf("user id param is invalid"))
		return
	}

	post, err := a.app.GetPostByID(ctx, postID.(int))
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Get post by id failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"success": true,
			"message": "Get post by id successful",
			"post":    post,
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Get post by id failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// DeletePostByID godoc
// @Summary Delete a post by ID
// @Description Remove a post from the system by its unique ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} map[string]interface{} "Post deletion successful"
// @Failure 400 {object} map[string]interface{} "Bad request error with message"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/posts/{id} [delete]
func (a *api) DeletePostByID(ctx *gin.Context) {
	postID, err := GetParamByName(ctx, "id")
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Delete post failed",
				"error":   fmt.Errorf("post id param is invalid").Error(),
			},
		}, fmt.Errorf("post id param is invalid"))
		return
	}

	err = a.app.DeletePostByID(ctx, postID.(int))
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Delete post failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"success": true,
			"message": "Delete post successful",
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Delete post failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// UpdatePostByID godoc
// @Summary Update a post by ID
// @Description Update the details of an existing post by its unique ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Param post body model.Post true "Updated post data"
// @Success 200 {object} map[string]interface{} "Post update successful"
// @Failure 400 {object} map[string]interface{} "Bad request error with message"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/posts/{id} [put]
func (a *api) UpdatePostByID(ctx *gin.Context) {
	postID, err := GetParamByName(ctx, "id")
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Update post failed",
				"error":   fmt.Errorf("post id param is invalid").Error(),
			},
		}, fmt.Errorf("post id param is invalid"))
		return
	}

	err = a.app.UpdatePostByID(ctx, postID.(int))
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Update post failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"success": true,
			"message": "Update post successful",
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Update post failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// LikePostByID godoc
// @Summary Like a post by ID
// @Description Like a specific post by its unique ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} map[string]interface{} "Like post successful"
// @Failure 400 {object} map[string]interface{} "Bad request error with message"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/posts/{id}/like [put]
func (a *api) LikePostByID(ctx *gin.Context) {
	err := a.app.LikePostByID(ctx)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Like post failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"success": true,
			"message": "Like post successful",
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Like post failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// UnLikePostByID godoc
// @Summary Unlike a post by ID
// @Description Unlike a specific post by its unique ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 200 {object} map[string]interface{} "Unlike post successful"
// @Failure 400 {object} map[string]interface{} "Bad request error with message"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/posts/{id}/unlike [put]
func (a *api) UnLikePostByID(ctx *gin.Context) {
	err := a.app.UnlikePostByID(ctx)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Unlike post failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"success": true,
			"message": "Unlike post successful",
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Unlike post failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// Login godoc
// @Summary Login a user and generate a JWT token
// @Description Logs in the user with their credentials and returns a JWT token for authentication
// @Tags auth
// @Produce json
// @Param loginInput body model.LoginInput true "Login credentials"
// @Success 200 {object} map[string]interface{} "Login successful with token and user data"
// @Failure 400 {object} map[string]interface{} "Bad request error with message"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/auth/login [post]
func (a *api) Login(ctx *gin.Context) {
	var loginInput model.LoginInput

	if err := ctx.ShouldBindJSON(&loginInput); err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Login failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	user, err := a.app.Login(ctx, loginInput)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Login failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	token, err := createJwtToken(user.Username)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Login failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"success": true,
			"message": "Login successful",
			"token":   token,
			"user":    user,
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Login failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}

// Register godoc
// @Summary Register a new user
// @Description Registers a new user with the provided credentials and returns user data
// @Tags auth
// @Produce json
// @Param registerInput body model.RegisterInput true "Register credentials"
// @Success 200 {object} map[string]interface{} "Register successful with user data"
// @Failure 400 {object} map[string]interface{} "Bad request error with message"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/auth/register [post]
func (a *api) Register(ctx *gin.Context) {
	var registerInput model.RegisterInput

	if err := ctx.ShouldBindJSON(&registerInput); err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Register failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	user, err := a.app.Register(ctx, registerInput)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusBadRequest, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Register failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}

	err = writeJSON(ctx.Writer, http.StatusOK, map[string]interface{}{
		"data": map[string]interface{}{
			"success": true,
			"message": "Register successful",
			"user":    user,
		},
	}, nil)
	if err != nil {
		serverErrorResponse(ctx.Writer, ctx.Request, http.StatusInternalServerError, map[string]interface{}{
			"data": map[string]interface{}{
				"success": false,
				"message": "Register failed",
				"error":   err.Error(),
			},
		}, err)
		return
	}
}
