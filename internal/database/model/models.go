package model

type Models struct {
	User User
	UserResponse UserResponse
	Post Post
}

func NewModels() Models {
	return Models{
		User: User{},
		UserResponse: UserResponse{},
		Post: Post{},
	}
}
