package types

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}
type GetUserResponse struct {
	*ApiResponse
	Users []*User
}

type CreateUserResponse struct {
	*ApiResponse
}

type UpdateUserResponse struct {
	*ApiResponse
}

type DeleteUserResponse struct {
	*ApiResponse
}
