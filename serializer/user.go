package serializer

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token,omitempty"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}
